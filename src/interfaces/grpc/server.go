package grpc

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/hamakn/go_grpc_sample/src/proto"
)

type Config struct {
	Address string
}

type server struct {
	grpcServer *grpc.Server
	config     *Config
}

func NewServer(config *Config) (*server, error) {
	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(s, NewGreeterServer())

	return &server{
		grpcServer: s,
		config:     config,
	}, nil
}

func (s *server) Serve() error {
	lis, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		return err
	}

	reflection.Register(s.grpcServer)

	return s.grpcServer.Serve(lis)
}

func (s *server) Shutdown() error {
	// TODO
	return nil
}
