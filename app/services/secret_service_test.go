package services

import (
	"testing"
)

func TestSecretServiceImpl_CreateSecret(t *testing.T) {
	data := []byte("data")

	service := NewSecretService()
	_, err := service.CreateSecret(data)
	if err != nil {
		t.Error(err)
	}
}
