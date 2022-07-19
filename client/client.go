package client

import (
	"context"
	"encoding/json"
	"github.com/r3labs/sse/v2"
	"sync"
)

const (
	DataEndpoint = "/relay/v1/data/events/realtime"
)

// SSEClient encapsulates sse.Client and stores the topic to listen to on the Flashbots Relayer.
type SSEClient struct {
	// Context management.
	ctx    context.Context
	cancel context.CancelFunc

	// Client management.
	wg     sync.WaitGroup
	client *sse.Client

	// Messages management.
	topics   []EventType
	messages chan *sse.Event
}

// New creates an SSE client listening to the SSE endpoint on the provided relay.
// It stores the topics to listen to later on.
func New(opts ...Option) (*SSEClient, error) {
	s := &settings{}

	// Apply all provided options.
	s.apply(append(defaultOptions(), opts...))

	// Verify applied options.
	if err := s.validate(); err != nil {
		return nil, err
	}

	// This way, we can cancel the child without cancelling the parent context.
	ctx, cancelFunc := context.WithCancel(s.ctx)
	return &SSEClient{
		ctx:      ctx,
		cancel:   cancelFunc,
		topics:   s.topics,
		client:   sse.NewClient(s.relayURL+DataEndpoint, s.opts...),
		messages: make(chan *sse.Event),
	}, nil
}

func (c *SSEClient) Subscribe(stream string) (chan Result, error) {
	// We use a channel that sends message with the associated error instead of two separate
	// channels as it is thread safe and avoid out-of-order responses.
	results := make(chan Result)

	if err := c.client.SubscribeChanWithContext(c.ctx, stream, c.messages); err != nil {
		return nil, err
	}

	c.wg.Add(1)

	go func() {
		defer c.wg.Done()

		for {
			select {
			case <-c.ctx.Done():
				return
			case message := <-c.messages:
				result := Result{
					Message: &Data{},
					Error:   nil,
				}

				result.Error = json.Unmarshal(message.Data, &result.Message)
				results <- result
			}
		}
	}()

	return results, nil
}

func (c *SSEClient) Unsubscribe() {
	c.client.Unsubscribe(c.messages)
	c.cancel()
	c.wg.Wait()
}
