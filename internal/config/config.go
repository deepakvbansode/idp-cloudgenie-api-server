package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)
type WebServerConfig struct {
	Port        string `required:"true"`
	RoutePrefix string `required:"false" split_words:"true" default:"/cloud-genie"`
	LogLevel    string `required:"false" split_words:"true" default:"info"`
	Github      GithubConfig
	Crossplane  CrossplaneConfig
	Mongo       MongoConfig
}


const (
	envPrefix      = "CG"
	configFileName = "../../etc/config/config.localhost.env"
)
func GetEnvConfig() (*WebServerConfig, error) {
	err := godotenv.Load(configFileName)
	if err != nil {
		fmt.Println("No config files found to load to env. Defaulting to environment", err)
	}

	config := &WebServerConfig{}

	err = envconfig.Process(envPrefix, config)
	if err != nil {
		fmt.Println("failed to process env variables:", err)
		return nil, err
	}
	return config, nil
}