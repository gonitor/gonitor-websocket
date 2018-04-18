package disk

import (
	"github.com/gonitor/gonitor-websocket/util"
	"github.com/gonitor/gonitor/service/disk"
)

// ServiceGetUsage gets the disk usage in bytes.
func ServiceGetUsage() []byte {
	result, err := disk.ServiceGetUsage()
	return util.ConvertInterfaceToJsonBytes(result, err)
}
