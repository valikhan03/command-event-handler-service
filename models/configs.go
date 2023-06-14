package models

import (
	"bytes"
	"encoding/json"
	"os"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type KafkaConfig struct {
	Brokers []string `json:"addr"`
	GroupID string   `json:"group_id"`
	Topics  []string `json:"topics"`
}

var KafkaConf KafkaConfig

func InitKafkaConfigs() {
	configjson, err := os.ReadFile("configs/consumer_group.json")
	if err != nil {
		panic(err)
	}
	config := KafkaConfig{}

	err = json.NewDecoder(bytes.NewBuffer(configjson)).Decode(&config)
	if err != nil {
		panic(err)
	}

	KafkaConf = config
}

type ElasticConfig struct {
	Addrs    []string `json:"addrs"`
	Username string   `json:"username"`
	Password string   `json:"-"`
}

var ElasticConf ElasticConfig



func InitElasticConfigs() {
	configjson, err := os.ReadFile("configs/elastic.json")
	if err != nil {
		panic(err)
	}
	config := ElasticConfig{}

	err = json.NewDecoder(bytes.NewBuffer(configjson)).Decode(&config)
	if err != nil {
		panic(err)
	}

	ElasticConf = config
}








var DBConfs	DBConfigs



type DBConfigs struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"-"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}




func InitDBConfigs() {
	dbConfFile, err := ioutil.ReadFile("configs/db.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(dbConfFile, &DBConfs)
	if err != nil{
		log.Fatal(err)
	}

	DBConfs.Password = os.Getenv("DB_PASSWORD")
}


type ServerConfigs struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}