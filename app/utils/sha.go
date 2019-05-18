package utils

import (
	"crypto/sha256"
	"crypto/sha512"
)

func ToSha512(b []byte) []byte{
	h := sha512.New()
	h.Write(b)
	return h.Sum(nil)
}

func ToSha256(b []byte) []byte{
	h := sha256.New()
	h.Write(b)
	return h.Sum(nil)
}