package sse

import (
	"context"
	"errors"
	testing2 "github.com/0xpanoramix/frd-go/internal/testing"
	"github.com/0xpanoramix/frd-go/topics"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
)

func TestSubscribeToValidBids(t *testing.T) {
	var wg sync.WaitGroup

	srv := testing2.NewMockServer(DataEndpoint)
	defer func() {
		err := srv.Close()
		assert.NoError(t, err)

		wg.Wait()
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := srv.Serve(); !errors.Is(err, http.ErrServerClosed) {
			assert.NoError(t, err)
		}
	}()

	ctx := context.Background()
	opts := []Option{
		WithRelay("http://127.0.0.1:8080"),
		WithTopics(topics.BuilderBidValid, topics.ProposerGetHeader),
		WithContext(ctx),
	}

	client, err := New(opts...)
	assert.NoError(t, err)

	res, err := client.Subscribe(testing2.MockStream)
	assert.NoError(t, err)

	data := <-res
	assert.Equal(t, data.Message.EventType, topics.BuilderBidValid)

	data = <-res
	assert.Equal(t, data.Message.EventType, topics.ProposerGetHeader)

	client.Unsubscribe()
}
