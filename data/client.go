package data

import (
	"encoding/json"
	"fmt"
	"github.com/flashbots/go-boost-utils/types"
	"net/http"
	"strings"
	"time"
)

// TransparencyClient is used to wrap calls to relays.
type TransparencyClient struct {
	clt     http.Client
	baseURL string
}

// NewTransparencyClient creates a new TransparencyClient using a relay URL and a request timeout.
func NewTransparencyClient(baseURL string, timeout time.Duration) *TransparencyClient {
	return &TransparencyClient{
		clt: http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}

type GetProposerPayloadsDeliveredOptions struct {
	Slot        uint64
	Cursor      uint64
	Limit       uint64
	BlockHash   string
	BlockNumber uint64
}

func (o *GetProposerPayloadsDeliveredOptions) ToQueryParameters() string {
	var args []string

	if o.Slot > 0 {
		args = append(args, fmt.Sprintf("slot=%d", o.Slot))
	}

	if o.Cursor > 0 {
		args = append(args, fmt.Sprintf("cursor=%d", o.Cursor))
	}

	if o.Limit > 0 {
		args = append(args, fmt.Sprintf("limit=%d", o.Limit))
	}

	if o.BlockHash != "" {
		args = append(args, fmt.Sprintf("block_hash=%s", o.BlockHash))
	}

	if o.BlockNumber > 0 {
		args = append(args, fmt.Sprintf("block_number=%d", o.BlockNumber))
	}

	params := strings.Join(args, "&")
	if len(params) > 0 {
		params = "?" + params
	}

	return params
}

// GetProposerPayloadsDelivered provides bids for payloads that were delivered to proposers.
// nolint:lll
func (c *TransparencyClient) GetProposerPayloadsDelivered(options *GetProposerPayloadsDeliveredOptions) ([]types.BidTrace, error) {
	// Creates the URL using the provided the base url and options.
	path := "/relay/v1/data/bidtraces/proposer_payload_delivered"
	url := c.baseURL + path

	if options != nil {
		url += options.ToQueryParameters()
	}

	res, err := http.Get(url) //nolint
	if err != nil {
		return nil, err
	}

	// Handle error in the response.
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("failed request to %s with status code %d", url, res.StatusCode)
	}

	// Extract bid traces from response.
	var traces []types.BidTrace
	if err = json.NewDecoder(res.Body).Decode(&traces); err != nil {
		return nil, err
	}

	return traces, nil
}
