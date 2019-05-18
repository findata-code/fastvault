package configurations

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Port          string `json:"port"`
	Key           string `json:"key"`
	SecretStorage string `json:"secret_storage"`
}

var Current Configuration

func Load(filepath string) error {
	log.Println("use", filepath)
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &Current)
	if err != nil {
		return err
	}

	return nil
}
