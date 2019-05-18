package app

import (
	"fastvault/app/configurations"
	"os"
)

func InitialConfig() {
	err := configurations.Load(os.Getenv("ENV_CONFIG_LOCATION"))
	if err != nil {
		panic(err)
	}
}