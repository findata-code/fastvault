package services

import (
	"fastvault/app/utils"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	FILE_SERVICE_FILE_EXTENSION = ".secret"
)

type FileService interface {
	CreateSecretFile(filename string, body []byte) error
	ReadSecretFile(filename string) ([]byte, error)
}

type FileServiceImpl struct {
	SecretDirectory string
	Key             string
}

func NewFileService(directory, key string) FileService {
	if len(key)%16 != 0 {
		return nil
	}

	return &FileServiceImpl{
		SecretDirectory: strings.TrimSuffix(directory, "/"),
		Key:             key,
	}
}

func (fs *FileServiceImpl) CreateSecretFile(token string, body []byte) error {
	eb, err := utils.CFBEncrypt([]byte(fs.Key), body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fs.getDictatoryString(token), eb, 0622)
}

func (fs *FileServiceImpl) ReadSecretFile(token string) ([]byte, error) {
	b, err := ioutil.ReadFile(fs.getDictatoryString(token))
	if err != nil {
		return nil, err
	}

	db, err := utils.CFBDecrypt([]byte(fs.Key), b)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (fs *FileServiceImpl) getDictatoryString(filename string) string {
	return fmt.Sprintf("%s/%s%s", fs.SecretDirectory, filename, FILE_SERVICE_FILE_EXTENSION)
}
