package secret

import (
	"fastvault/app/configurations"
	"fastvault/app/services"
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
