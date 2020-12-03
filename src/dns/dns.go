package dns

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config test something
type Config struct {
	Name   string `yaml:"name"`
	String string `yaml:"test"`
}

// ImportConfigs test
func (config *Config) ImportConfigs() error {

	filename, _ := filepath.Abs("./test.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &config)
	return err
}
