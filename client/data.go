package client

import "github.com/0xpanoramix/frd-go/dto"

type Data struct {
	Timestamp string       `json:"timestamp"`
	EventType EventType    `json:"event_type"`
	EventData dto.BidTrace `json:"event_data"`
}
