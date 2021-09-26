package consumer

import (
	"context"
	"github.com/lubedd/broker-bible/internal/connect"
	pb "github.com/lubedd/broker-bible/internal/proto"
	"google.golang.org/grpc"
)

type Consumer interface {
	Send(routerKey, message string, id uint64) error
	Recv() (*pb.Consumer, error)
}

type consumerProcess struct {
	conn   *grpc.ClientConn
	client pb.BrokerClient
	stream pb.Broker_ConsumerChatClient
}

func NewConsumer(ip, port string) (*consumerProcess, error) {
	conn, err := connect.NewConnection(ip, port)
	if err != nil {
		return nil, err
	}

	client := connect.NewClient(conn)

	stream, err := client.ConsumerChat(context.Background())
	if err != nil {
		return nil, err
	}

	return &consumerProcess{
		conn:   conn,
		client: connect.NewClient(conn),
		stream: stream,
	}, nil
}

func (c *consumerProcess) Send(routerKey, message string, id uint64) error {
	return c.stream.Send(&pb.Consumer{RoutingKey: routerKey, Message: message, Id: id})
}
func (c *consumerProcess) Recv() (*pb.Consumer, error) {
	return c.stream.Recv()
}
