package main

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

type consumer interface {
	Messages() <-chan *sarama.ConsumerMessage
	Close() error
}

type partCon struct {
	consumer sarama.PartitionConsumer
}

func (pc *partCon) Messages() <-chan *sarama.ConsumerMessage {
	return pc.consumer.Messages()
}

func (pc *partCon) Close() error {
	pc.consumer.AsyncClose()
	return nil
}

type clusterCon struct {
	consumer *cluster.Consumer
}

func (pc *clusterCon) Messages() <-chan *sarama.ConsumerMessage {
	return pc.consumer.Messages()
}

func (pc *clusterCon) Close() error {
	return pc.consumer.Close()
}
