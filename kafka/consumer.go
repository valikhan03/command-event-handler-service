package kafka

import (
	"log"

	"command-event-handler-service/models"
	"command-event-handler-service/services"
	"encoding/json"

	"github.com/Shopify/sarama"
)

func InitConsumerGroup() sarama.ConsumerGroup {
	configConsumer := sarama.NewConfig()
	configConsumer.Consumer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup(models.KafkaConf.Brokers, models.KafkaConf.GroupID, configConsumer)
	if err != nil {
		panic(err)
	}

	return consumerGroup
}

type handler struct {
	eventHandlerChan chan<- *services.Event
}

func InitConsumerHandler(eventsChan chan *services.Event) *handler {
	return &handler{
		eventHandlerChan: eventsChan,
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (h *handler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (h *handler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, 
								claim sarama.ConsumerGroupClaim) error {
	for {
		msg := <-claim.Messages()
		var event services.Event
		err := json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("Consume Claims error: %s\n", err.Error())
		}
		log.Printf("Received event %s\n", event.Command)

		h.eventHandlerChan <- &event

		session.MarkMessage(msg, "")
	}
}
