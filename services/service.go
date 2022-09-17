package services

import (
	"bytes"
	"command-event-handler-service/elastic"
	"command-event-handler-service/models"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Event models.Event

func (e *Event) CreateAuction() {
	data, err := json.Marshal(e)
	if err != nil{

	}

	req := esapi.CreateRequest{
		Index: "auctions",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil{

	}

	if res.IsError() {

	}
}

func (e *Event) UpdateAuction() {
	data, err := json.Marshal(e)
	if err != nil{

	}

	req := esapi.UpdateRequest{
		Index: "auctions",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil{

	}

	if res.IsError(){

	}
}

func (e *Event) CancelAuction() {
	data, err := json.Marshal(e)
	if err != nil{

	}

	req := esapi.UpdateRequest{
		Index: "auctions",
		DocumentID: e.Entity["id"].(string),
		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil{

	}

	if res.IsError(){

	}
}

func (e *Event) AddParticipant() {
	data, err := json.Marshal(e)
	if err != nil{

	}

	req := esapi.CreateRequest{
		Index: "auction-participants",
		DocumentID: e.Entity["id"].(string),
		
		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil{

	}

	if res.IsError(){

	}
}

func (e *Event) DeleteParticipant() {
	req := esapi.DeleteRequest{
		Index: "auction-participants",
		DocumentID: e.Entity["id"].(string),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil{

	}

	if res.IsError(){

	}
}	

func (e *Event) AddProduct() {
	data, err := json.Marshal(e)
	if err != nil{

	}

	req := esapi.CreateRequest{
		Index: "auction-products",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil{

	}

	if res.IsError(){

	}
}

func (e *Event) UpdateProduct() {
	data, err := json.Marshal(e)
	if err != nil{

	}

	req := esapi.UpdateRequest{
		Index: "auction-products",
		DocumentID: e.Entity["id"].(string),

		Body: bytes.NewReader(data),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil{

	}

	if res.IsError(){

	}
}

func (e *Event) DeleteProduct() {
	req := esapi.DeleteRequest{
		Index: "auction-products",
		DocumentID: e.Entity["id"].(string),
	}

	res, err := req.Do(context.Background(), elastic.ElasticConn)
	if err != nil{

	}

	if res.IsError(){

	}
}
