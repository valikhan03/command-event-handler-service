package services

import (
	"bytes"
	"command-event-handler-service/elastic"
	"command-event-handler-service/models"
	"context"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Event models.Event

func (e *Event) CreateAuction() {
	data, err := json.Marshal(e)
	if err != nil {
		log.Printf("Event parse error: %x\n", err)
	}

	req := esapi.CreateRequest{
		Index:      "auctions",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %x\n", err)
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %x\n", res.String())
	}
}

func (e *Event) UpdateAuction() {
	data, err := json.Marshal(e)
	if err != nil {
		log.Printf("Event parse error: %x\n", err)
	}

	req := esapi.UpdateRequest{
		Index:      "auctions",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %x\n", err)
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %x\n", res.String())
	}
}

func (e *Event) DeleteAuction() {
	data, err := json.Marshal(e)
	if err != nil {
		log.Printf("Event parse error: %x\n", err)
	}

	req := esapi.UpdateRequest{
		Index:      "auctions",
		DocumentID: e.Entity["id"].(string),
		Body:       bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %x\n", err)
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %x\n", res.String())
	}
}

func (e *Event) AddParticipant() {
	data, err := json.Marshal(e)
	if err != nil {
		log.Printf("Event parse error: %x\n", err)
	}

	req := esapi.CreateRequest{
		Index:      "auction-participants",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %x\n", err)
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %x\n", res.String())
	}
}

func (e *Event) DeleteParticipant() {
	req := esapi.DeleteRequest{
		Index:      "auction-participants",
		DocumentID: e.Entity["id"].(string),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %x\n", err)
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %x\n", res.String())
	}
}

func (e *Event) AddLot() {
	data, err := json.Marshal(e)
	if err != nil {
		log.Printf("Event parse error: %x\n", err)
	}

	req := esapi.CreateRequest{
		Index:      "auction-products",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %x\n", err)
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %x\n", res.String())
	}
}

func (e *Event) UpdateLot() {
	data, err := json.Marshal(e)
	if err != nil {
		log.Printf("Event parse error: %x\n", err)
	}

	req := esapi.UpdateRequest{
		Index:      "auction-products",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %x\n", err)
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %x\n", res.String())
	}
}

func (e *Event) DeleteLot() {
	req := esapi.DeleteRequest{
		Index:      "auction-products",
		DocumentID: e.Entity["id"].(string),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil {
		log.Printf("Elasticsearch request error: %x\n", err)
	}

	if res.IsError() {
		log.Printf("Elasticsearch request error: %x\n", res.String())
	}
}
