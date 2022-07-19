package dto

import (
	"github.com/0xpanoramix/frd-go/topics"
	"github.com/flashbots/go-boost-utils/types"
)

type Result struct {
	Message *Data
	Error   error
}

type Data struct {
	Timestamp string           `json:"timestamp"`
	EventType topics.EventType `json:"event_type"`
	EventData types.BidTrace   `json:"event_data"`
}
