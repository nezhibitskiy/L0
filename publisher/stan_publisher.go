package main

import (
	"L0/internal"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

const (
	clusterID = "test-cluster"
	clientID  = "delivery-service-publisher"
)

// publishEvent publish an event via NATS Streaming server
func publishEvent(data *internal.Order) {
	// Connect to NATS Streaming server
	sc, err := stan.Connect(
		clusterID,
		clientID,
		stan.NatsURL("nats://127.0.0.1:4223"),
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer sc.Close()
	channel := "order-channel"
	byteData, _ := json.Marshal(data)
	eventMsg := byteData

	err = sc.Publish(channel, eventMsg)
	if err != nil {
		return
	}
	log.Println("Published message on channel: " + channel)
}

func main() {
	order := internal.Order{
		OrderUid:    "2",
		TrackNumber: "",
		Entry:       "",
		Delivery: struct {
			Name    string `json:"name"`
			Phone   string `json:"phone"`
			Zip     string `json:"zip"`
			City    string `json:"city"`
			Address string `json:"address"`
			Region  string `json:"region"`
			Email   string `json:"email"`
		}{},
		Payment: struct {
			Transaction  string    `json:"transaction"`
			RequestId    string    `json:"request_id"`
			Currency     string    `json:"currency"`
			Provider     string    `json:"provider"`
			Amount       int       `json:"amount"`
			PaymentDt    time.Time `json:"payment_dt"`
			Bank         string    `json:"bank"`
			DeliveryCost int       `json:"delivery_cost"`
			GoodsTotal   int       `json:"goods_total"`
			CustomFee    int       `json:"custom_fee"`
		}{},
		Items:             nil,
		Locale:            "",
		InternalSignature: "",
		CustomerId:        "",
		DeliveryService:   "",
		Shardkey:          "",
		SmId:              0,
		DateCreated:       time.Time{},
		OofShard:          "",
	}
	publishEvent(&order)
}
