package utils

import (
	"encoding/json"
	"os"
)

func WriteJSON(path string, data interface{}) error {
	bytes,_ := json.MarshalIndent(data,"", " ")
	return os.WriteFile(path, bytes, 0644)
}

func ReadJSON(path string, target interface{}) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, target)
}