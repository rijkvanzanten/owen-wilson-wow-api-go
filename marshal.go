package main

import (
	"bytes"
	"encoding/json"
)

// The golang std encoding/json Marshal escapes <, >, and & for "convenience" This function is a
// drop-in replacement for json.Marshal(), but it won't escape the aforementioned characters
func Marshal(i interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder((buffer))
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(i)
	return bytes.TrimRight(buffer.Bytes(), "\n"), err
}