package app

import (
	"fastvault/app/configurations"
	"os"
)

var Config configurations.Configuration

func InitialConfig() {
	c, err := configurations.Load(os.Getenv("ENV_CONFIG_LOCATION"))
	if err != nil {
		c =  configurations.Configuration {
			Port: ":8080",
		}
	}

	Config = c
}