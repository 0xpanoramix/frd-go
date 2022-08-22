package main

import (
	"github.com/0xpanoramix/frd-go/data"
	"log"
	"time"
)

func main() {
	clt := data.NewTransparencyClient("https://builder-relay-ropsten.flashbots.net/", time.Second)

	traces, err := clt.GetProposerPayloadsDelivered(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(traces)
}
