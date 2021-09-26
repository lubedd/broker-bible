package producer

import (
	"broker/bible/internal/connect"
	pb "broker/bible/internal/proto"
	"context"
	"google.golang.org/grpc"
)

type Producer interface {
	NewTask(routingKey, messageText string) (*pb.ResponseProducer, error)
}

type ProducerProcess struct {
	conn   *grpc.ClientConn
	client pb.BrokerClient
}

func NewProducer(ip, port string) (*ProducerProcess, error) {
	conn, err := connect.NewConnection(ip, port)
	if err != nil {
		return nil, err
	}

	return &ProducerProcess{
		conn:   conn,
		client: connect.NewClient(conn),
	}, nil
}

func (p *ProducerProcess) NewTask(routingKey, messageText string) (*pb.ResponseProducer, error) {
	request := &pb.RequestProducer{
		MessageText: messageText,
		RoutingKey:  routingKey,
	}

	return p.client.AddMessage(context.Background(), request)
}
