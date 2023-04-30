package services

import (
	"reflect"
	"github.com/valikhan03/tool"
)

type EventHandler struct {
	eventHandlerChan <-chan *Event
}

func InitEventHandler(eventsChan chan *Event) *EventHandler {
	return &EventHandler{
		eventHandlerChan: eventsChan,
	}
}

// run as a goroutine
func (e *EventHandler) HandleCommandEvents() {
	for {
		eventObj := <-e.eventHandlerChan
		method := tool.EventFuncs[eventObj.Command]
		reflect.ValueOf(eventObj).MethodByName(method).Call([]reflect.Value{})
	}
}

