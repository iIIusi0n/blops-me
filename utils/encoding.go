package utils

import (
	"encoding/base64"
)

func EncodeBase(input string) string {
	return base64.URLEncoding.EncodeToString([]byte(input))
}

func DecodeBase(encoded string) string {
	decoded, _ := base64.URLEncoding.DecodeString(encoded)
	return string(decoded)
}
