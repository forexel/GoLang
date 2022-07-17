package main

import (
	"fmt"
    
    "GOLANG-1/Lesson8/configparse"
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

func main() {
	appConfig := &AppConfig{}
	// Load default config file
    (configparse.AppConfig).ReadYaml("conf.yaml")
	// Load specific configuration with env variables
	(configparse.AppConfig).ReadEnv()
	fmt.Printf("%+v\n", *appConfig)
	// go run src/main.go
	// {Name:test Port:8444 Db:{User:admin Password:1234 Host:localhost Port:5432}}
}
