package load

import (
	"github.com/gonitor/gonitor-websocket/util"
	"github.com/gonitor/gonitor/service/load"
)

// ServiceGetAverage gets the load average in bytes.
func ServiceGetAverage() []byte {
	result, err := load.ServiceGetAverage()
	return util.ConvertInterfaceToJsonBytes(result, err)
}

// ServiceGetMisc gets the load misc in bytes.
func ServiceGetMisc() []byte {
	result, err := load.GetMisc()
	return util.ConvertInterfaceToJsonBytes(result, err)
}
