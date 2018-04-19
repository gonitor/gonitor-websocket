package host_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor-websocket/service/host"
)

// TestServiceGetInfo .
func TestServiceGetInfo(test *testing.T) {
	result := host.ServiceGetInfo()
	assert.Equal(test, result != nil, true)
}

// TestServiceGetTemperature .
func TestServiceGetTemperature(test *testing.T) {
	result := host.ServiceGetTemperature()
	assert.Equal(test, result != nil, true)
}
