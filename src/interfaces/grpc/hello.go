package grpc

import (
	"context"

	pb "github.com/hamakn/go_grpc_sample/src/proto"
)

type greeterServerImpl struct {
}

func NewGreeterServer() *greeterServerImpl {
	return &greeterServerImpl{}
}

func (g *greeterServerImpl) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
