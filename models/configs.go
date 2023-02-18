package models

import(
	"os"
	"bytes"
	"encoding/json"
)

type KafkaConfig struct {
	Brokers []string `json:"addr"`
	GroupID string   `json:"group_id"`
	Topics  []string `json:"topics"`
}

var KafkaConf KafkaConfig

func InitKafkaConfigs() {
	configjson, err := os.ReadFile("configs/consumer_group.json")
	if err != nil{
		panic(err)
	}
	config := KafkaConfig{}

	err = json.NewDecoder(bytes.NewBuffer(configjson)).Decode(&config)
	if err != nil{
		panic(err)
	}

	KafkaConf = config
}