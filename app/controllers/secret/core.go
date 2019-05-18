package secret

import (
	"fastvault/app/configurations"
	"fastvault/app/services"
	"fastvault/app/utils"
	"fmt"
	"log"
)

var secretService services.SecretService
var fileService services.FileService

func Initialise() {
	log.Println("controllers/secret initialised")

	secretService = services.NewSecretService()
	fileService = services.NewFileService(
		configurations.Current.SecretStorage,
		configurations.Current.Key)
}

func filename(token []byte) string {
	return fmt.Sprintf("%x", utils.ToSha512([]byte(token)))
}