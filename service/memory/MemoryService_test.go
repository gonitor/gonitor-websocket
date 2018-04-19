package memory_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor-websocket/service/memory"
)

// TestServiceGetVirtualMemory .
func TestServiceGetVirtualMemory(test *testing.T) {
	result := memory.ServiceGetVirtualMemory()
	assert.Equal(test, result != nil, true)
}

// TestServiceGetSwapMemory .
func TestServiceGetSwapMemory(test *testing.T) {
	result := memory.ServiceGetSwapMemory()
	assert.Equal(test, result != nil, true)
}
