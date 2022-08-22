# frd-go

1. [Introduction](#introduction)
2. [How do the clients work ?](#how-do-the-clients-work-)
    1. [HTTP Client](#http-client)
    2. [SSE Client](#sse-client)
3. [Getting Started !](#getting-started-)
    1. [Installation](#installation)
    2. [Quickstart](#quickstart)
4. [Contributing](#contributing)
5. [Author](#author)

## Introduction

frd-go is a package which lets you use several clients to interact with Flashbots Relays,
specifically with the [Flashbots Data Transparency API](https://flashbots.notion.site/Relay-API-Spec-5fb0819366954962bc02e81cb33840f5#38a21c8a40e64970904500eb7b373ea5):

- An HTTP client for standard requests to the relays.
- A client for SSE (server-sent events) events sent by the relays.

## How do the clients work ?

### HTTP Client
Using the http package from the goland standard library, it makes requests to the relay you've
provided.
List of supported endpoints:
- [x] `/relay/v1/data/bidtraces/proposer_payload_delivered`

### SSE Client
Using the [sse package](https://github.com/r3labs/sse) made by r3labs, it connects to the
provided relayer and subscribe to incoming events: `BidTrace`.

These bids are forwarded (along with an optional error that may have happened during parsing) in a
channel you can use in your own application.

For now, it only supports connection to a single relay.

## Getting started !

### Installation

```shell
go get github.com/0xpanoramix/frd-go
```

### Quickstart

Below is an example of how you can create a client and make requests to the relay:
```go
package main

import (
	"github.com/0xpanoramix/frd-go/constants"
	"github.com/0xpanoramix/frd-go/data"
	"log"
	"time"
)

func main() {
	clt := data.NewTransparencyClient(constants.FlashbotsRelayRopsten, time.Second)

	traces, err := clt.GetProposerPayloadsDelivered(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(traces)
}
```

Below is an example of how you can create a sse-client:
```go
package main

import (
	"context"
	"fmt"
	"github.com/0xpanoramix/frd-go/sse"
	"github.com/0xpanoramix/frd-go/topics"
	"log"
)

func main() {
	opts := []sse.Option{
		sse.WithRelay("http://127.0.0.1:8080"),
		sse.WithTopics(topics.BuilderBidValid),
		sse.WithContext(context.Background()),
	}
	clt, err := sse.New(opts...)
	if err != nil {
		log.Fatal(err)
	}

	res, err := clt.Subscribe("messages")

	data := <-res
	// This will print a Flashbots BidTrace.
	fmt.Println(data.Message.EventData)

	clt.Unsubscribe()
}

```

## Contributing

Make sure you have Go installed
```shell
go version
```

Then to build the project:
```shell
go build
```

And to run the tests:
```shell
go test ./... -v -race
```

## Author

Made with ❤️ by 🤖 [Luca Georges François](https://github.com/0xpanoramix) 🤖