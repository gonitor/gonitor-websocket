package disk_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor-websocket/service/disk"
)

// TestServiceGetUsage .
func TestServiceGetUsage(test *testing.T) {
	result := disk.ServiceGetUsage()
	assert.Equal(test, result != nil, true)
}
