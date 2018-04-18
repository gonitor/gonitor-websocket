package env

import (
	"os"
	"strconv"
)

// EnableCompression .
var EnableCompression = true

// ReadBufferSize .
var ReadBufferSize = 1024

// WriteBufferSize .
var WriteBufferSize = 1024

// LoadEnvVariables loads environment vaiables.
func LoadEnvVariables() {
	if os.Getenv("GONI_COMPRESS") == "false" {
		EnableCompression = false
	}

	readBufferSizeEnv, readBufferSizeEnvErr := strconv.Atoi(os.Getenv("GONI_RBSIZE"))
	if readBufferSizeEnvErr == nil && readBufferSizeEnv > 0 {
		ReadBufferSize = readBufferSizeEnv
	}

	writeBufferSizeEnv, writeBufferSizeEnvErr := strconv.Atoi(os.Getenv("GONI_WBSIZE"))
	if writeBufferSizeEnvErr == nil && writeBufferSizeEnv > 0 {
		WriteBufferSize = writeBufferSizeEnv
	}
}
