package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor-websocket/handler"
)

// WebSocketV1EndPoint .
var WebSocketV1EndPoint = "/websocket/v1"

// WebSocketV1SetRoutes .
func WebSocketV1SetRoutes(router *gin.Engine) {
	router.GET(WebSocketV1EndPoint, handler.HanldeWebSocket)
}
