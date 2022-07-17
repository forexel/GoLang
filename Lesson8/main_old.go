package main
import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "log"
)
type conf struct {
	port         []string `yaml:"port", envconfig:"SERVER_PORT"`
	db_url       []string `yaml:"db_url", envconfig:"DB_URL"`
	jaeger_url   []string `yaml:"jaeger_url", envconfig:"JAEGER_URL"`
	sentry_url   []string `yaml:"sentry_url", envconfig:"SENTY_URL"`
	kafka_broker []string `yaml:"kafka_broker", envconfig:"KAFKA_BROKER"`
	some_app_id  []string `yaml:"some_app_id", envconfig:"SOME_APP_ID"`
	some_app_key []string `yaml:"some_app_key", envconfig:"SOME_APP_KEY"`
}
func (c *conf) getConf() *conf {
    yamlFile, err := ioutil.ReadFile("conf.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    return c
}
func main() {
    var c conf
    c.getConf()
    fmt.Println(c)
}