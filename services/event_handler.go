package services

import (
	"command-event-handler-service/models"
	"reflect"
	"github.com/valikhan03/tool"
)

type EventHandler struct {
	eventHandlerChan <-chan *models.Event
}

func InitEventHandler(eventsChan chan *models.Event) *EventHandler {
	return &EventHandler{
		eventHandlerChan: eventsChan,
	}
}

// run as a goroutine
func (e *EventHandler) HandleCommandEvents() {
	for {
		eventObj := <-e.eventHandlerChan
		method := tool.EventFuncs[eventObj.Command]
		reflect.ValueOf(Event(*eventObj)).MethodByName(method)
	}
}