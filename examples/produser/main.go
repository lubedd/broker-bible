package main

import (
	"github.com/lubedd/broker-bible"
	"github.com/lubedd/broker-bible/pkg/domain"
	"log"
)

func main() {
	producer, err := bible.GetProducerClient("127.0.0.1", "5300")
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	response, err := producer.NewTask("a", "Very important message")
	if err != nil || response.Message != domain.ProducerTaskAccepted {
		log.Fatalf("can not create new task %v", err)
	}
}
