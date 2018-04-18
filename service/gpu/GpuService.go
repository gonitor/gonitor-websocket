package gpu

import (
	"github.com/gonitor/gonitor-websocket/util"
	"github.com/gonitor/gonitor/service/gpu"
)

// ServiceGetInfo gets the GPU information in bytes.
func ServiceGetInfo() []byte {
	result, err := gpu.ServiceGetInfo()
	return util.ConvertInterfaceToJsonBytes(result, err)
}
