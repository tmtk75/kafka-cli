package main

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit an offset",
	Run: func(cmd *cobra.Command, args []string) {
		Commit()
	},
}

func init() {
	RootCmd.AddCommand(commitCmd)
	//pflags := commitCmd.PersistentFlags()
}

func Commit() {
	err := initProfile()
	if err != nil {
		log.Fatal(err)
	}

	dst, err := NewDestination()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", dst)

	b := sarama.NewBroker(dst.Hosts[0])

	cfg, err := NewConfig(true)
	if err != nil {
		log.Fatal(err)
	}

	err = b.Open(cfg)
	if err != nil {
		log.Fatal(err)
	}

	req := sarama.OffsetCommitRequest{
		Version:                 1, // 0.8.2 or later
		ConsumerGroup:           dst.Group,
		ConsumerGroupGeneration: sarama.GroupGenerationUndefined,
	}

	req.AddBlock(dst.Topic, dst.Partition, dst.Offset, time.Now().Unix(), "")
	a, s, err := req.Offset(dst.Topic, dst.Partition)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("a: %v, s: %v", a, s)

	res, err := b.CommitOffset(&req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", res)
}
