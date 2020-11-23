package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DbHostname string `json:"DB_HOSTNAME"`
	DbUser     string `json:"DB_USER"`
	DbPassword string `json:"DB_PASSWORD"`
	DbName     string `json:"DB_NAME"`
}

func GetConfig(params ...string) Configuration {

	configuration := Configuration{}

	env := "dev"

	if len(params) > 0 {

		env = params[0]

	}

	fileName := fmt.Sprintf("./config/%s_config.json", env)

	gonfig.GetConf(fileName, &configuration)

	return configuration
}
