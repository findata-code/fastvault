package services

import (
	"fmt"
	"os"
	"testing"
)

const (
	TEST_DATA = "data"
	TEST_KEY = "0123456789abcdef"
	TEST_DIRECTORY = "."
	TEST_FILE = "test01"
	TEST_FULL_PATH = "./test01.secret"
	TEST_TOKEN = "test01"
)

func TestFileServiceImpl_CreateSecretFile(t *testing.T) {
	defer func() {
		os.Remove(TEST_FULL_PATH)
	}()

	service := NewFileService(TEST_DIRECTORY, TEST_KEY)
	err := service.CreateSecretFile(TEST_FILE, []byte(TEST_DATA))
	if err != nil {
		t.Error(err)
	}

	_, err = os.Open(TEST_FULL_PATH)
	if err != nil {
		t.Error(err)
	}
}

func TestFileServiceImpl_ReadSecretFile(t *testing.T) {
	defer func() {
		os.Remove(TEST_FULL_PATH)
	}()

	service := NewFileService(TEST_DIRECTORY, TEST_KEY)
	err := service.CreateSecretFile(TEST_FILE, []byte(TEST_DATA))
	if err != nil {
		t.Error(err)
	}

	_, err = os.Open(TEST_FULL_PATH)
	if err != nil {
		t.Error(err)
	}

	b, err := service.ReadSecretFile(TEST_FILE)
	if err != nil {
		t.Error(err)
	}

	if fmt.Sprintf("%s", b) != TEST_DATA {
		t.Error("test data is not correct")
	}
}