package testing

import (
	"encoding/json"
	"github.com/0xpanoramix/frd-go/dto"
	"github.com/0xpanoramix/frd-go/topics"
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

func NewMockServer(pattern string) *MockServer {
	server := &MockServer{sseSRV: sse.New()}

	server.sseSRV.CreateStream(MockStream)

	server.mux = http.NewServeMux()
	server.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		for _, topic := range r.URL.Query()["topics"] {
			server.Publish(topics.EventType(topic))
		}
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

func (s *MockServer) Publish(topic topics.EventType) {
	var data []byte

	switch topic {
	case topics.BuilderBidValid:
		data, _ = json.Marshal(dto.Data{
			EventType: topics.BuilderBidValid,
		})
	case topics.ProposerGetHeader:
		data, _ = json.Marshal(dto.Data{
			EventType: topics.ProposerGetHeader,
		})
	case topics.ProposerSubmitBlindedBlock:
		data, _ = json.Marshal(dto.Data{
			EventType: topics.ProposerSubmitBlindedBlock,
		})
	}

	s.sseSRV.Publish(MockStream, &sse.Event{
		Data:  data,
		Event: []byte(topic),
	})
}

func (s *MockServer) Close() error {
	s.sseSRV.Close()

	return s.srv.Close()
}
