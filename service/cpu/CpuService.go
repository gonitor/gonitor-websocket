package cpu

import (
	"github.com/gonitor/gonitor-websocket/util"
	"github.com/gonitor/gonitor/service/cpu"
)

// ServiceGetSummaryPercent in bytes.
func ServiceGetSummaryPercent() []byte {
	result, err := cpu.ServiceGetSummaryPercent()
	return util.ConvertInterfaceToJsonBytes(result, err)
}

// ServiceGetCount gets the CPU counts in bytes.
func ServiceGetCount() []byte {
	result, err := cpu.ServiceGetCount()
	return util.ConvertInterfaceToJsonBytes(result, err)
}

// ServiceGetSummaryTime gets the time summary in bytes.
func ServiceGetSummaryTime() []byte {
	result, err := cpu.ServiceGetSummaryTime()
	return util.ConvertInterfaceToJsonBytes(result, err)
}

// ServiceGetInfo gets the CPU information in bytes.
func ServiceGetInfo() []byte {
	result, err := cpu.ServiceGetInfo()
	return util.ConvertInterfaceToJsonBytes(result, err)
}

// ServiceGetPercent gets CPU usage in bytes.
func ServiceGetPercent() []byte {
	result, err := cpu.ServiceGetPercent()
	return util.ConvertInterfaceToJsonBytes(result, err)
}

// ServiceGetTime gets CPU time information in bytes.
func ServiceGetTime() []byte {
	result, err := cpu.ServiceGetTime()
	return util.ConvertInterfaceToJsonBytes(result, err)
}
