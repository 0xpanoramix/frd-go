package client

import (
	"context"
	"github.com/r3labs/sse/v2"
	"net/url"
)

type Option func(*settings)

// WithRelay is used to provide the relay to connect to.
func WithRelay(relayURL string) Option {
	return func(s *settings) {
		s.relayURL = relayURL
	}
}

// WithTopics is used to provide which topics to listen to.
func WithTopics(topics ...EventType) Option {
	return func(s *settings) {
		s.topics = topics
	}
}

// WithSSEClientOptions is used to customize SSE Client.
func WithSSEClientOptions(opts ...func(c *sse.Client)) Option {
	return func(s *settings) {
		s.opts = opts
	}
}

// WithContext is used to provide a specific external context to the client subscription.
func WithContext(ctx context.Context) Option {
	return func(s *settings) {
		s.ctx = ctx
	}
}

// defaultOptions is used to provide default options for the context and topics to subscribe to.
func defaultOptions() []Option {
	return []Option{
		WithContext(context.Background()),
		WithTopics(BuilderBidValid, ProposerGetHeader, ProposerSubmitBlindedBlock),
	}
}

type settings struct {
	ctx      context.Context
	relayURL string
	topics   []EventType
	opts     []func(c *sse.Client)
}

func (s *settings) apply(opts []Option) {
	for _, opt := range opts {
		opt(s)
	}
}

func (s *settings) validate() error {
	if _, err := url.ParseRequestURI(s.relayURL); err != nil {
		return err
	}

	return nil
}
