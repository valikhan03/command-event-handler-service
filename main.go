package main

import (
	"command-event-handler-service/elastic"
	"command-event-handler-service/kafka"
	"command-event-handler-service/models"
	"command-event-handler-service/services"
	"context"
	"log"
)

func main() {
	models.InitElasticConfigs()
	elastic.InitElasticConn()

	eventsChan := make(chan *services.Event)
	
	models.InitKafkaConfigs()
	consumerGroup := kafka.InitConsumerGroup()
	consumerHandler := kafka.InitConsumerHandler(eventsChan)

	eventHandler := services.InitEventHandler(eventsChan)

	go 	eventHandler.HandleCommandEvents()
	log.Println("command-handler started...")

	for {
		err := consumerGroup.Consume(context.Background(), models.KafkaConf.Topics, consumerHandler)
		if err != nil {
			log.Fatalf("Kafka Consumer error: %s", err.Error())
		}
	}
}
