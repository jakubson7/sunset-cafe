package utils

import "encoding/json"

func ToJSON(v any) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}
func ToBytesJSON(v any) ([]byte, error) {
	return json.Marshal(v)
}
