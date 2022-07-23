package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Port         string `yaml:"port", envconfig:"SERVER_PORT"`
	Db_url       string `yaml:"db_url", envconfig:"DB_URL"`
	Jaeger_url   string `yaml:"jaeger_url", envconfig:"JAEGER_URL"`
	Sentry_url   string `yaml:"sentry_url", envconfig:"SENTY_URL"`
	Kafka_broker string `yaml:"kafka_broker", envconfig:"KAFKA_BROKER"`
	Some_app_id  string `yaml:"some_app_id", envconfig:"SOME_APP_ID"`
	Some_app_key string `yaml:"some_app_key", envconfig:"SOME_APP_KEY"`
}

func (c *AppConfig) getConf() *AppConfig {

	yamlFile, err := ioutil.ReadFile("conf.yaml")

	
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	_, err = url.Parse(c.Db_url)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	_, err = url.Parse(c.Jaeger_url)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	_, err = url.Parse(c.Sentry_url)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c AppConfig
	c.getConf()
	fmt.Println(c.Port)
}
