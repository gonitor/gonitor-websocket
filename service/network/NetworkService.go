package network

import (
	"github.com/gonitor/gonitor-websocket/util"
	"github.com/gonitor/gonitor/service/network"
)

// ServiceGetInterfaces gets the network interaces in bytes.
func ServiceGetInterfaces() []byte {
	result, err := network.ServiceGetInterfaces()
	return util.ConvertInterfaceToJsonBytes(result, err)
}

// ServiceGetConnections gets the network connections in bytes.
func ServiceGetConnections() []byte {
	result, err := network.ServiceGetConnections()
	return util.ConvertInterfaceToJsonBytes(result, err)
}
