package runtime

import (
	"github.com/gonitor/gonitor/service/runtime"

	"github.com/gonitor/gonitor-websocket/util"
)

// ServiceGetGoOS gets the operating system (OS) of Go runtime in bytes.
func ServiceGetGoOS() []byte {
	result := runtime.ServiceGetGoOS()
	return util.ConvertInterfaceToJsonBytes(result, nil)
}
