package handler

import (
	"encoding/json"
)

func jsonError(message string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}

	bytes, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}
	return bytes
}
