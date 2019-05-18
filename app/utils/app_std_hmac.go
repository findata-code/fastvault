package utils

import (
	"crypto/hmac"
	"crypto/sha256"
)

func ToAppStdHmac(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}