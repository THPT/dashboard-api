package main

import (
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	configPath = "./config/config.yaml"
)

var config configuration

type dbConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type configuration struct {
	Default dbConfig `yaml:"default"`
}

func init() {
	var fileName string
	var yamlFile []byte
	var err error

	if fileName, err = filepath.Abs(configPath); err != nil {
		panic(err)
	}

	if yamlFile, err = ioutil.ReadFile(fileName); err != nil {
		panic(err)
	}

	config = configuration{}
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		panic(err)
	}
}

func openConnection() (*gorm.DB, error) {
	conf := config.Default
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)
	return gorm.Open(conf.Type, connectionString)
}
