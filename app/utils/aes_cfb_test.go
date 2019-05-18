package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCFBEncrypt(t *testing.T) {
	key := []byte("0123456789abcdef")
	data := []byte("data")

	ciphertext, err := CFBEncrypt(key, data)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%x\n", ciphertext)

	decryptedData, err := CFBDecrypt(key, ciphertext)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(data, decryptedData) {
		t.Error("data and decrypted data should be equal")
	}
}