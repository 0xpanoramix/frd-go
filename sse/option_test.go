package sse

import (
	"github.com/0xpanoramix/frd-go/topics"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplyAndValidateSettings(t *testing.T) {
	testCases := []struct {
		name string
		opts []Option

		expectedError    bool
		expectedSettings *settings
	}{
		{
			name: "Valid options",
			opts: []Option{
				WithRelay("http://localhost:8080"),
				WithTopics(topics.BuilderBidValid, topics.ProposerGetHeader, topics.ProposerSubmitBlindedBlock),
			},
			expectedError: false,
			expectedSettings: &settings{
				relayURL: "http://localhost:8080",
				topics: []topics.EventType{
					topics.BuilderBidValid,
					topics.ProposerGetHeader,
					topics.ProposerSubmitBlindedBlock,
				},
				opts: nil,
			},
		},
		{
			name: "Invalid relay URL",
			opts: []Option{
				WithRelay("d"),
				WithTopics(topics.BuilderBidValid, topics.ProposerGetHeader),
			},
			expectedError: true,
			expectedSettings: &settings{
				relayURL: "d",
				topics: []topics.EventType{
					topics.BuilderBidValid,
					topics.ProposerGetHeader,
				},
				opts: nil,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := &settings{}

			// Apply options.
			s.apply(tt.opts)

			// Verify provided options.
			err := s.validate()

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedSettings, s)
		})
	}
}
