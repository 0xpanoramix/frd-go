package data

import (
	"github.com/0xpanoramix/frd-go/constants"
	"github.com/flashbots/go-boost-utils/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewTransparencyClient(t *testing.T) {
	client := NewTransparencyClient(constants.FlashbotsRelayMainnet, time.Second)

	assert.NotNil(t, client)
}

func TestGetProposerPayloadsDelivered(t *testing.T) {
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
			name:          "Valid request on Mainnet",
			baseURL:       constants.FlashbotsRelayMainnet,
			options:       nil,
			expectedError: false,
		},
		{
			name:          "Valid request on Sepolia",
			baseURL:       constants.FlashbotsRelaySepolia,
			options:       nil,
			expectedError: false,
		},
		{
			name:          "Valid request on Goerli",
			baseURL:       constants.FlashbotsRelayGoerli,
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

func TestGetValidatorRegistration(t *testing.T) {
	publicKey := types.PublicKey{}
	//nolint:lll
	err := publicKey.UnmarshalText([]byte("0xb606e206c2bf3b78f53ebff8be8e8d4af2f0da68646b5642c4d511b15ab5ddb122ae57b48eab614f8ca5bafbe75a5999"))
	require.NoError(t, err)

	testCases := []struct {
		name          string
		baseURL       string
		publicKey     types.PublicKey
		expectedError bool
	}{
		{
			name:          "Get valid validator registration",
			baseURL:       constants.FlashbotsRelayMainnet,
			publicKey:     publicKey,
			expectedError: false,
		},
		{
			name:          "Get invalid validator registration",
			baseURL:       constants.FlashbotsRelayMainnet,
			publicKey:     types.PublicKey{},
			expectedError: true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			clt := NewTransparencyClient(tt.baseURL, time.Second)

			registration, err := clt.GetValidatorRegistration(tt.publicKey)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, registration)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, registration)
			}
		})
	}
}

func TestGetBuilderBlocksReceived(t *testing.T) {
	testCases := []struct {
		name          string
		baseURL       string
		options       *GetBuilderBlocksReceivedOptions
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
			name:          "Valid request on Mainnet",
			baseURL:       constants.FlashbotsRelayMainnet,
			options:       nil,
			expectedError: false,
		},
		{
			name:    "Valid request on Mainnet with limit",
			baseURL: constants.FlashbotsRelayMainnet,
			options: &GetBuilderBlocksReceivedOptions{
				Limit: 5,
			},
			expectedError: false,
		},
		{
			name:          "Valid request on Sepolia",
			baseURL:       constants.FlashbotsRelaySepolia,
			options:       nil,
			expectedError: false,
		},
		{
			name:          "Valid request on Goerli",
			baseURL:       constants.FlashbotsRelayGoerli,
			options:       nil,
			expectedError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			clt := NewTransparencyClient(tt.baseURL, time.Second)

			traces, err := clt.GetBuilderBlocksReceived(tt.options)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, traces)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, traces)

				if tt.options != nil && tt.options.Limit > 0 {
					assert.Len(t, traces, int(tt.options.Limit))
				}
			}
		})
	}
}
