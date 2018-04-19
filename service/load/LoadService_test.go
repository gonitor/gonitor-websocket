package load_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor-websocket/service/load"
)

// TestServiceGetAverage .
func TestServiceGetAverage(test *testing.T) {
	result := load.ServiceGetAverage()
	assert.Equal(test, result != nil, true)
}

// TestServiceGetMisc .
func TestServiceGetMisc(test *testing.T) {
	result := load.ServiceGetMisc()
	assert.Equal(test, result != nil, true)
}
