package network_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor-websocket/service/network"
)

// TestServiceGetInterfaces .
func TestServiceGetInterfaces(test *testing.T) {
	result := network.ServiceGetInterfaces()
	assert.Equal(test, result != nil, true)
}
