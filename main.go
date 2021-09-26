package bible

import (
	"github.com/lubedd/broker-bible/pkg/consumer"
	"github.com/lubedd/broker-bible/pkg/producer"
)
//TODO: Скрыть *pb.Consumer и *pb.ResponseProducer под уровнем абстракции

func GetProducerClient(ip, port string) (producer.Producer, error) {
	return producer.NewProducer(ip, port)
}

func GetConsumerClient(ip, port string) (consumer.Consumer, error) {
	return consumer.NewConsumer(ip, port)
}
