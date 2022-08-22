package data

import (
	"github.com/0xpanoramix/frd-go/constants"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTransparencyClient(t *testing.T) {
	client := NewTransparencyClient(constants.FlashbotsRelayRopsten, time.Second)

	assert.NotNil(t, client)
}
