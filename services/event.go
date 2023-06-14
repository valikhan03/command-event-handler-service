package services

import (
	"bytes"
	"command-event-handler-service/elastic"
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"strconv"

	database "command-event-handler-service/db"

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
	var attempt_reqs tool.AttemptRequirements
	data, err := json.Marshal(e.Entity)
	if err != nil {
		log.Printf("Event parse error: %s\n", err.Error())
	}
	query := `select id, auction_id, approve_required, enter_fee_required, enter_fee_amount 
			  from tb_attempt_requirements where auction_id=$1`
	database.DBConn.Get(&attempt_reqs, query, e.Entity["auction_id"].(string))

	if attempt_reqs.EnterFee > 0 {
		query := `select id, auction_id, approve_required, enter_fee_required, enter_fee_amount 
			  from tb_attempt_requirements where auction_id=$1`
		database.DBConn.Get(&attempt_reqs, query, e.Entity["auction_id"].(string))
	}
	req := esapi.CreateRequest{
		Index:      "auction-participants",
		DocumentID: e.Entity["auction_id"].(string),

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

func (e *Event) CreateInvoice() {
	product_name := getProductName(e.Entity["product_type"].(int), e.Entity["product_id"].(int))

	tx, err := database.DBConn.BeginTxx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		log.Printf("DB error: %v\n", err)
		return
	}
	query := `insert into tb_invoices
			  (user_id, product_type, product_id, product_name, amount, currency, p_date_insert)
			  values ($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.Exec(query, e.Entity["user_id"].(int), e.Entity["product_type"].(int), e.Entity["product_id"].(int), product_name, e.Entity["price"].(int))
	if err != nil {
		log.Printf("DB error: %v", err)
		tx.Rollback()
	}

	tx.Commit()
}