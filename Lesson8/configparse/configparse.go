package configparse

import (
	"fmt"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	port         string `yaml:"port", envconfig:"SERVER_PORT"`
	db_url       string `yaml:"db_url", envconfig:"DB_URL"`
	jaeger_url   string `yaml:"jaeger_url", envconfig:"JAEGER_URL"`
	sentry_url   string `yaml:"sentry_url", envconfig:"SENTY_URL"`
	kafka_broker string `yaml:"kafka_broker", envconfig:"KAFKA_BROKER"`
	some_app_id  string `yaml:"some_app_id", envconfig:"SOME_APP_ID"`
	some_app_key string `yaml:"some_app_key", envconfig:"SOME_APP_KEY"`
}

func (config *AppConfig) ReadEnv() {
	envconfig.Process("app", config)
}

func (config *AppConfig) ReadYaml(filepath string) {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error Reading Config file with path: %v\n", filepath)
	}
	yaml.Unmarshal(yamlFile, config)
	fmt.Println(config.port)
}