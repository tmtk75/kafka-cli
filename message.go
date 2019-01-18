package main

import (
	"time"

	"github.com/Shopify/sarama"
)

type Message struct {
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	Topic     string    `json:"topic"`
	Partition int32     `json:"partition"`
	Offset    int64     `json:"offset"`
	Timestamp time.Time `json:"timestamp"` // only set if kafka is version 0.10+
}

func (m *Message) fill(c *sarama.ConsumerMessage) {
	m.Key = string(c.Key)
	m.Value = string(c.Value)
	m.Topic = c.Topic
	m.Partition = c.Partition
	m.Offset = c.Offset
	m.Timestamp = c.Timestamp
}
