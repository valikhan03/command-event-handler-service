package kafka

import (
	"command-event-handler-service/models"
	"encoding/json"

	"github.com/Shopify/sarama"
)

func InitConsumerGroup() sarama.ConsumerGroup {
	configConsumer := sarama.Config{}
	configConsumer.Consumer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup(models.KafkaConf.Brokers, models.KafkaConf.GroupID, &configConsumer)
	if err != nil{
		panic(err)
	}
	

	return consumerGroup
}

type handler struct{
	eventHandlerChan chan <- *models.Event
}

func InitConsumerHandler(eventsChan chan *models.Event) *handler{
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
func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for{
		msg := <- claim.Messages()
		var event models.Event
		err := json.Unmarshal(msg.Value, &event)
		if err != nil{

		}

		h.eventHandlerChan <- &event		

		session.MarkMessage(msg, "")
	}
}