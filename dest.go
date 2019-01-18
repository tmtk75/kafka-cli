package main

import "encoding/json"

type Destination struct {
	TLSInsecure bool
	Hosts       []string
	Topic       string
	Partition   int32
	Group       string
	Offset      int64
}

func (d *Destination) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func NewDestination() (*Destination, error) {
	var d Destination
	d.TLSInsecure = subv.GetBool(KeyTLSInsecure)
	d.Hosts = subv.GetStringSlice(KeyHosts)
	d.Topic = eitherString(flagTopic, subv.GetString(KeyTopic))
	d.Partition = eitherInt32(flagPartition, subv.GetInt32(KeyPartition))
	d.Group = eitherString(flagGroup, subv.GetString(KeyGroup))

	_offset := eitherString(flagOffset, subv.GetString(KeyOffset))
	offset, err := ParseOffset(_offset)
	if err != nil {
		return nil, err
	}
	d.Offset = offset

	return &d, nil
}
