package route

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor-websocket/config"
	"github.com/gonitor/gonitor/route"
)

// SetRoutes .
func SetRoutes(router *gin.Engine) {
	WebSocketV1SetRoutes(router)
	if config.EnableRestAPI {
		route.RestV1SetRoutes(router)
	}
	if config.EnableStreamAPI {
		route.StreamV1SetRoutes(router)
	}
}

// GetWebSocketEndPoint .
func GetWebSocketEndPoint(url string) string {
	root := WebSocketV1GroupEndPoint

	var buffer bytes.Buffer
	buffer.WriteString(root)
	buffer.WriteString(url)
	return buffer.String()
}
