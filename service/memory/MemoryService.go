package memory

import (
	"github.com/gonitor/gonitor-websocket/util"
	"github.com/gonitor/gonitor/service/memory"
)

// ServiceGetVirtualMemory gets the virtual memory in bytes.
func ServiceGetVirtualMemory() []byte {
	result, err := memory.ServiceGetVirtualMemory()
	return util.ConvertInterfaceToJsonBytes(result, err)
}

// ServiceGetSwapMemory gets the swap memory in bytes.
func ServiceGetSwapMemory() []byte {
	result, err := memory.ServiceGetSwapMemory()
	return util.ConvertInterfaceToJsonBytes(result, err)
}
