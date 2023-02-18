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
	elastic.InitElasticConn()

	eventsChan := make(chan *models.Event)

	consumerGroup := kafka.InitConsumerGroup()
	consumerHandler := kafka.InitConsumerHandler(eventsChan)

	eventHandler := services.InitEventHandler(eventsChan)

	go func() {
		for {
			err := consumerGroup.Consume(context.Background(), models.KafkaConf.Topics, consumerHandler)
			if err != nil {
				log.Fatal("Kafka Consumer error^")
			}
		}
	}()

	go func() {
		eventHandler.HandleCommandEvents()
	}()
}
