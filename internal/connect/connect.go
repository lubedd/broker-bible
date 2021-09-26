package connect

import (
	pb "github.com/lubedd/broker-bible/internal/proto"
	"fmt"
	"google.golang.org/grpc"
)



func NewConnection(ip, port string) (*grpc.ClientConn, error) {
	return grpc.Dial(fmt.Sprintf("%s:%s", ip, port),grpc.WithInsecure())
}

func NewClient(conn *grpc.ClientConn) pb.BrokerClient {
	return pb.NewBrokerClient(conn)
}
