package host

import (
	"github.com/gonitor/gonitor-websocket/util"
	"github.com/gonitor/gonitor/service/host"
)

// ServiceGetInfo gets the host information in bytes.
func ServiceGetInfo() []byte {
	result, err := host.ServiceGetInfo()
	return util.ConvertInterfaceToJsonBytes(result, err)
}

// ServiceGetTemperature gets the host temperature in bytes.
func ServiceGetTemperature() []byte {
	result, err := host.ServiceGetTemperature()
	return util.ConvertInterfaceToJsonBytes(result, err)
}
