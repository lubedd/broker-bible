package main

import (
	"github.com/lubedd/broker-bible"
	"github.com/lubedd/broker-bible/pkg/domain"
	"log"
	"time"
)

func main() {
	consumer, err := bible.GetConsumerClient("127.0.0.1", "5300")
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	var id uint64
	err = consumer.Send("a", domain.ConsumersOpenConnection, id)
	if err != nil {
		log.Fatalf("can not send %v", err)
	}

	for {
		resp, err := consumer.Recv()
		if err != nil {
			continue
		}
		id = resp.Id
		DoSomething(resp.Message)

		err = consumer.Send(resp.RoutingKey, domain.ConsumersTaskAccepted, id)
		if err != nil {
			continue
		}
	}
}

func DoSomething(message string) {
	log.Println(message)
	time.Sleep(1 * time.Second)
}
