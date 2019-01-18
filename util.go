package main

import (
	"strconv"

	"github.com/Shopify/sarama"
)

const (
	OffsetOldest = sarama.OffsetOldest
	OffsetNewest = sarama.OffsetNewest
	Oldest       = "oldest"
	Newest       = "newest"
)

func ParseOffset(offset string) (int64, error) {
	switch offset {
	case Oldest:
		return OffsetOldest, nil
	case Newest:
		return OffsetNewest, nil
	}
	return strconv.ParseInt(offset, 10, 64)
}

func eitherString(a *string, b string) string {
	if a != nil && *a != "" {
		return *a
	}
	return b
}

func eitherInt32(a *int32, b int32) int32 {
	if a != nil && *a != 0 {
		return *a
	}
	return b
}
