package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/spf13/cobra"
)

var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consume a topic",
	Run: func(cmd *cobra.Command, args []string) {
		Consume()
	},
}

func init() {
	RootCmd.AddCommand(consumeCmd)

	pflags := consumeCmd.PersistentFlags()
	flagWithoutGroup = pflags.Bool("without-group", false, "Consume without consumer group if it's given")
}

func Consume() {
	err := initProfile()
	if err != nil {
		log.Fatal(err)
	}

	dst, err := NewDestination()
	if err != nil {
		log.Fatal(err)
	}
	logger.Printf("%v", dst)

	ConsumePartition := func() consumer {
		cfg, err := NewConfig(dst.TLSInsecure)
		if err != nil {
			log.Fatal(err)
		}

		con, err := sarama.NewConsumer(dst.Hosts, cfg)
		if err != nil {
			log.Fatal(err)
		}

		pc, err := con.ConsumePartition(dst.Topic, dst.Partition, dst.Offset)
		if err != nil {
			log.Fatal(err)
		}
		logger.Printf("consumer started")

		return &partCon{consumer: pc}
	}

	ConsumeCluster := func() consumer {
		cfg, err := NewConfigCluster(dst.TLSInsecure)
		if err != nil {
			log.Fatal(err)
		}

		con, err := cluster.NewConsumer(dst.Hosts, dst.Group, []string{dst.Topic}, cfg)
		if err != nil {
			log.Fatal(err)
		}

		return &clusterCon{consumer: con}
	}

	var con consumer
	if dst.Group != "" && !*flagWithoutGroup {
		logger.Printf("in cluster mode. the given offset is ignored. offset: %v", dst.Offset)
		con = ConsumeCluster()
	} else {
		logger.Printf("in partition mode.")
		con = ConsumePartition()
	}
	defer con.Close()

	go func() {
		for msg := range con.Messages() {
			m := Message{}
			m.fill(msg)
			b, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", string(b))
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)

	select {
	case <-s:
		//pc.AsyncClose()
	}

	if err := con.Close(); err != nil {
		log.Printf("failed to close. %v", err)
	}
}
