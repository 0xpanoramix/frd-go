package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetProposerPayloadsDelivered(t *testing.T) {
	ropstenBaseURL := "https://builder-relay-ropsten.flashbots.net/"
	testCases := []struct {
		name          string
		baseURL       string
		options       *GetProposerPayloadsDeliveredOptions
		expectedError bool
	}{
		{
			name:          "Invalid base URL",
			baseURL:       "",
			options:       nil,
			expectedError: true,
		},
		/*
			{
				name:    "Invalid block hash",
				baseURL: ropstenBaseURL,
				options: &GetProposerPayloadsDeliveredOptions{
					BlockHash: "",
				},
				expectedError: true,
			},
		*/
		{
			name:          "Valid request",
			baseURL:       ropstenBaseURL,
			options:       nil,
			expectedError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			clt := NewTransparencyClient(tt.baseURL, time.Second)

			traces, err := clt.GetProposerPayloadsDelivered(tt.options)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, traces)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, traces)
			}
		})
	}
}
