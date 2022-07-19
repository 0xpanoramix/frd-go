package client

import (
	"github.com/flashbots/go-boost-utils/types"
)

type Data struct {
	Timestamp string         `json:"timestamp"`
	EventType EventType      `json:"event_type"`
	EventData types.BidTrace `json:"event_data"`
}
