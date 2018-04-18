package util

import "encoding/json"

// ConvertInterfaceToJsonBytes converts any interface to json bytes.
func ConvertInterfaceToJsonBytes(input interface{}, inputError error) []byte {
	if inputError != nil {
		return nil
	}
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil
	}
	return bytes
}
