package client

import (
	"github.com/r3labs/sse/v2"
)

const (
	DataEndpoint = "/relay/v1/data/events/realtime"
)

// SSEClient encapsulates sse.Client and stores the topic to listen to on the Flashbots Relayer.
type SSEClient struct {
	topics []EventType
	client *sse.Client
}

// New creates an SSE client listening to the SSE endpoint on the provided relay.
// It stores the topics to listen to later on.
func New(opts ...Option) (*SSEClient, error) {
	s := &settings{}

	// Apply all provided options.
	s.apply(opts)

	// Verify applied options.
	if err := s.validate(); err != nil {
		return nil, err
	}

	return &SSEClient{
		topics: s.topics,
		client: sse.NewClient(s.relayURL+DataEndpoint, s.opts...),
	}, nil
}

func (c *SSEClient) Subscribe(stream string) (chan *sse.Event, error) {
	res := make(chan *sse.Event)

	err := c.client.SubscribeChan(stream, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
