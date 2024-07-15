package utils

import "encoding/base64"

func EncodeStringBase64(data string) string {
	encodedString := base64.StdEncoding.EncodeToString([]byte(data))
	return encodedString
}

func DecodeStringBase64(data string) string {
	decodedString, _ := base64.StdEncoding.DecodeString(data)
	return string(decodedString)
}
