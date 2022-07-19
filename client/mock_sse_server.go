package client

import (
	"encoding/json"
	"github.com/r3labs/sse/v2"
	"net/http"
	"time"
)

const (
	MockStream = "messages"
)

type MockServer struct {
	srv    *http.Server
	sseSRV *sse.Server
	mux    *http.ServeMux
}

func NewMockServer(topic EventType) *MockServer {
	server := &MockServer{sseSRV: sse.New()}

	server.sseSRV.CreateStream(MockStream)

	server.mux = http.NewServeMux()
	server.mux.HandleFunc(DataEndpoint, func(w http.ResponseWriter, r *http.Request) {
		server.Publish(topic)
		server.sseSRV.ServeHTTP(w, r)
	})

	server.srv = &http.Server{
		ReadHeaderTimeout: time.Second,
		Addr:              ":8080",
		Handler:           server.mux,
	}

	return server
}

func (s *MockServer) Serve() error {
	return s.srv.ListenAndServe()
}

func (s *MockServer) Publish(topic EventType) {
	switch topic {
	case BuilderBidValid:
		data, _ := json.Marshal(Data{
			EventType: BuilderBidValid,
		})

		s.sseSRV.Publish(MockStream, &sse.Event{
			Data:  data,
			Event: []byte(BuilderBidValid),
		})
	}
}

func (s *MockServer) Close() error {
	return s.srv.Close()
}
