package config

type CrossplaneConfig struct {
	APIEndpoint string `required:"true" split_words:"true" default:"http://localhost:8080"`
	APIToken    string `required:"true" split_words:"true"`
}
