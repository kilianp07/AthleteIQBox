package utils

import (
	"encoding/json"
	"os"
)

// ReadJSONFile reads the content of a JSON file at the specified path and unmarshals it into the target interface.
// The path parameter specifies the path of the JSON file to be read.
// The target parameter is a pointer to the interface into which the JSON content will be unmarshaled.
// Returns an error if there was a problem reading the file or unmarshaling the JSON content.
func ReadJSONFile(path string, target interface{}) (err error) {
	var (
		content []byte
	)

	content, err = os.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(content, target)
	return
}
