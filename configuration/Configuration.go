package configuration

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Configurationins AppConfigs

type AppConfigs struct {
	Development struct {
		Mode string `yaml:"mode"`
		Port string `yaml:"port"`
	}
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Sslmode  string `yaml:"sslmode"`
		Dbname   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	}
}

func Readconfigurationfile() error {

	file, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		return err
	}
	if yaml.Unmarshal(file, &Configurationins); err != nil {
		return err
	}
	return nil
}
