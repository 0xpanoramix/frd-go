package client

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
)

func TestSubscribeToValidBids(t *testing.T) {
	var wg sync.WaitGroup

	srv := NewMockServer(BuilderBidValid)
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

	client, err := New(WithRelay("http://127.0.0.1:8080"), WithTopics(BuilderBidValid))
	assert.NoError(t, err)

	res, err := client.Subscribe(MockStream)
	assert.NoError(t, err)

	data := <-res
	assert.Equal(t, EventType(data.Event), BuilderBidValid)
}
