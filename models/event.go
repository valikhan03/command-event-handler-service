package models

type Event struct {
	Command string                 `json:"command"`
	Entity  map[string]interface{} `json:"entity"`
}


