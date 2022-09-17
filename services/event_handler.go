package services

import (
	"command-event-handler-service/models"
	"encoding/json"
	"os"
	"reflect"
)

type EventHandler struct {
	eventHandlerChan <-chan *models.Event
	methods          map[string]string
}

func InitEventHandler(eventsChan chan *models.Event) *EventHandler {
	var m map[string]string
	file, err := os.ReadFile("configs/event-handler.json")
	if err != nil{
		
	}
	
	err = json.Unmarshal(file, &m)
	if err != nil{
		
	}
	
	return &EventHandler{
		eventHandlerChan: eventsChan,
		methods: m,
	}
}

// run as a goroutine
func (e *EventHandler) HandleCommandEvents() {
	for {
		eventObj := <-e.eventHandlerChan
		method := e.methods[eventObj.Command]
		reflect.ValueOf(Event(*eventObj)).MethodByName(method)
	}
}