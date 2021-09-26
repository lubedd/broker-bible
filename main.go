package bible

import (
	"github.com/lubedd/broker-bible/pkg/consumer"
	"github.com/lubedd/broker-bible/pkg/producer"
)

func GetProducerClient(ip, port string) (producer.Producer, error) {
	return producer.NewProducer(ip, port)
}

func GetConsumerClient(ip, port string) (consumer.Consumer, error) {
	return consumer.NewConsumer(ip, port)
}
