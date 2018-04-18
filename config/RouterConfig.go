package config

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor-websocket/route"
	goniconfig "github.com/gonitor/gonitor/config"
)

// SetRoutes sets all routes.
func SetRoutes(router *gin.Engine) {
	goniconfig.SetRoutes(router)
	route.WebSocketV1SetRoutes(router)
}

// GetWebSocketEndPoint concatenates the websocket endoint with a url.
func GetWebSocketEndPoint(url string) string {
	root := route.WebSocketV1EndPoint

	var buffer bytes.Buffer
	buffer.WriteString(root)
	buffer.WriteString(url)
	return buffer.String()
}
