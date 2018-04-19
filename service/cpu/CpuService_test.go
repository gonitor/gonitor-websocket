package cpu_test

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gonitor/gonitor-websocket/service/cpu"
)

// TestServiceGetSummaryPercent .
func TestServiceGetSummaryPercent(test *testing.T) {
	result := cpu.ServiceGetSummaryPercent()
	assert.Equal(test, result != nil, true)
}

// TestServiceGetCount .
func TestServiceGetCount(test *testing.T) {
	result := cpu.ServiceGetCount()
	assert.Equal(test, result != nil, true)
}

// TestServiceGetSummaryTime .
func TestServiceGetSummaryTime(test *testing.T) {
	result := cpu.ServiceGetSummaryTime()
	assert.Equal(test, result != nil, true)
}

// TestServiceGetInfo .
func TestServiceGetInfo(test *testing.T) {
	result := cpu.ServiceGetInfo()
	assert.Equal(test, result != nil, true)
}

// TestServiceGetPercent .
func TestServiceGetPercent(test *testing.T) {
	result := cpu.ServiceGetPercent()
	assert.Equal(test, result != nil, true)
}

// TestServiceGetTime .
func TestServiceGetTime(test *testing.T) {
	result := cpu.ServiceGetTime()
	assert.Equal(test, result != nil, true)
}
