package jetstream

import "time"

type AckType int

const (
	AckTypeAuto AckType = iota
	AckTypeManual
)

var defaultAckWait = 5 * time.Second
var defaultMaxRedeluver = 5

type ConsumerConfig struct {
	msgFilter     []string
	groupName    string
	ackType      AckType
	ackWait      time.Duration
	maxRedeliver int
}

type Option func(*ConsumerConfig)

func MsgFilter(msgFilter []string) Option {
	return func(cc *ConsumerConfig) {
		cc.msgFilter = msgFilter
	}
}


