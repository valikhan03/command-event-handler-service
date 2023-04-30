package models

import (
	"bytes"
	"encoding/json"
	"os"
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