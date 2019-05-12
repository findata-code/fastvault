package configurations

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Port string
}

func Load(filepath string) (Configuration, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = json.Unmarshal(b, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}
