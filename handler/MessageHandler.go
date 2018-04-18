package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonitor/gonitor-websocket/env"
	"github.com/gonitor/gonitor-websocket/service/cpu"
	"github.com/gonitor/gonitor-websocket/service/disk"
	"github.com/gonitor/gonitor-websocket/service/gpu"
	"github.com/gonitor/gonitor-websocket/service/host"
	"github.com/gonitor/gonitor-websocket/service/load"
	"github.com/gonitor/gonitor-websocket/service/memory"
	"github.com/gonitor/gonitor-websocket/service/network"
	"github.com/gonitor/gonitor-websocket/service/runtime"
	"github.com/gorilla/websocket"
)

// HanldeWebSocket handles websocket connection.
func HanldeWebSocket(context *gin.Context) {
	wsupgrader := GetUgrader()
	conn, err := wsupgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		return
	}
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		HandleMessage(conn, messageType, message)
	}
}

// GetUgrader .
func GetUgrader() websocket.Upgrader {
	return websocket.Upgrader{
		EnableCompression: env.EnableCompression,
		ReadBufferSize:    env.ReadBufferSize,
		WriteBufferSize:   env.WriteBufferSize,
	}
}

// HandleMessage handles message from client.
func HandleMessage(conn *websocket.Conn, messageType int, message []byte) {
	var response []byte
	messageString := string(message)
	switch messageString {
	// CPU
	case "CpuGetSummaryPercent":
		response = cpu.ServiceGetSummaryPercent()
	case "CpuGetCount":
		response = cpu.ServiceGetCount()
	case "CpuGetSummaryTime":
		response = cpu.ServiceGetSummaryTime()
	case "CpuGetInfo":
		response = cpu.ServiceGetInfo()
	case "CpuGetPercent":
		response = cpu.ServiceGetPercent()
	case "CpuGetTime":
		response = cpu.ServiceGetTime()

	// Disk
	case "DiskGetUsage":
		response = disk.ServiceGetUsage()

	// GPU
	case "GpuGetInfo":
		response = gpu.ServiceGetInfo()

	// Host
	case "HostGetInfo":
		response = host.ServiceGetInfo()
	case "HostGetTemperature":
		response = host.ServiceGetTemperature()

	// Load
	case "LoadGetAverage":
		response = load.ServiceGetAverage()
	case "LoadGetMisc":
		response = load.ServiceGetMisc()

	// Memory
	case "MemoryGetVirtualMemory":
		response = memory.ServiceGetVirtualMemory()
	case "MemoryGetSwapMemory":
		response = memory.ServiceGetSwapMemory()

	// Network
	case "NetworkGetInterfaces":
		response = network.ServiceGetInterfaces()
	case "NetworkGetConnections":
		response = network.ServiceGetConnections()

	// Runtime
	case "RuntimeGetGoOS":
		response = runtime.ServiceGetGoOS()
	}
	go func() {
		count := 0
		for {
			count++
			time.Sleep(1 * time.Second)
			conn.WriteMessage(messageType, response)
			if count >= 6 {
				break
			}
		}
	}()
}
