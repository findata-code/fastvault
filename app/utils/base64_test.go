package utils

import (
	"reflect"
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	data := []byte("base64")
	result := EncodeBase64(data)
	decoded, err := DecodeBase64(result)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(data, decoded) {
		t.Error("expect", data, "actual", decoded)
	}
}