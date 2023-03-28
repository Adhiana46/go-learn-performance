package main

import (
	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

func stdEncode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func stdDecode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func jsoniterEncode(data interface{}) ([]byte, error) {
	jsonI := jsoniter.ConfigCompatibleWithStandardLibrary
	return jsonI.Marshal(data)
}

func jsoniterDecode(data []byte, v interface{}) error {
	jsonI := jsoniter.ConfigCompatibleWithStandardLibrary
	return jsonI.Unmarshal(data, v)
}

func main() {
	// go test -bench=. -count 5 -benchmem
}
