package main

import (
	"gitlab.antyron.com/ITStudWay2017/location-service/service"
	"gitlab.antyron.com/ITStudWay2017/location-service/server"
	"github.com/ianschenck/envflag"
)

// Configuration options using environment variables
var (
	address = envflag.String("ADDRESS", ":8000", "")
	broker  = envflag.String("KAFKA_BROKER", "localhost:9092", "Elasticsearch index")
	retries = envflag.Int("KAFKA_PRODUCER_MAX_RETRY", 5, "Elasticsearch index")
	topic   = envflag.String("KAFKA_TOPIC", "location", "Elasticsearch index")
)

func main() {
	envflag.Parse()
	brokers := []string{*broker}
	messageBroker, err := service.NewKafkaBroker(brokers, *retries, *topic)

	if err != nil {
		panic(err)
	}
	srv := server.NewServer(*address, messageBroker)
	srv.Handle()
}
