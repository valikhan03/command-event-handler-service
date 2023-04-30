package services

import (
	"bytes"
	"command-event-handler-service/elastic"
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/valikhan03/tool"
)

type Event struct {
	Command string                 `json:"command"`
	Entity  map[string]interface{} `json:"entity"`
}

func (e *Event) CreateAuction() {
	data, err := json.Marshal(e.Entity)
	if err != nil {
		log.Printf("Event parse error: %s\n", err.Error())
	}

	req := esapi.CreateRequest{
		Index:      tool.AuctionsIDX,
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %s\n", err.Error())
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %s\n", res.String())
	}
}

func (e *Event) UpdateAuction() {
	data, err := json.Marshal(e.Entity)
	if err != nil {
		log.Printf("Event parse error: %s\n", err.Error())
	}

	req := esapi.UpdateRequest{
		Index:      tool.AuctionsIDX,
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %s\n", err.Error())
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %s\n", res.String())
	}
}

func (e *Event) DeleteAuction() {
	req := esapi.DeleteRequest{
		Index:      tool.AuctionsIDX,
		DocumentID: e.Entity["id"].(string),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %s\n", err.Error())
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %s\n", res.String())
	}
}

func (e *Event) AddParticipant() {
	data, err := json.Marshal(e.Entity)
	if err != nil {
		log.Printf("Event parse error: %s\n", err.Error())
	}

	req := esapi.CreateRequest{
		Index:      "auction-participants",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %s\n", err.Error())
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %s\n", res.String())
	}
}

func (e *Event) DeleteParticipant() {
	req := esapi.DeleteRequest{
		Index:      "auction-participants",
		DocumentID: e.Entity["id"].(string),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %s\n", err.Error())
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %s\n", res.String())
	}
}

func (e *Event) AddLot() {
	data, err := json.Marshal(e.Entity)
	if err != nil {
		log.Printf("Event parse error: %s\n", err.Error())
	}
	id := strconv.Itoa(int(e.Entity["id"].(float64)))
	req := esapi.CreateRequest{
		Index:      tool.LotsIDX,
		DocumentID: id,

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %s\n", err.Error())
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %s\n", res.String())
	}
}

func (e *Event) UpdateLot() {
	data, err := json.Marshal(e.Entity)
	if err != nil {
		log.Printf("Event parse error: %s\n", err.Error())
	}

	req := esapi.UpdateRequest{
		Index:      tool.LotsIDX,
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %s\n", err.Error())
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %s\n", res.String())
	}
}

func (e *Event) DeleteLot() {
	req := esapi.DeleteRequest{
		Index:      tool.LotsIDX,
		DocumentID: e.Entity["id"].(string),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %s\n", err.Error())
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %s\n", res.String())
	}
}
