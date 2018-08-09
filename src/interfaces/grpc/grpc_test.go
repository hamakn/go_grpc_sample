package grpc_test

import (
	"context"
	"testing"

	"github.com/hamakn/go_grpc_sample/src/interfaces/grpc"
	pb "github.com/hamakn/go_grpc_sample/src/proto"

	"github.com/stretchr/testify/require"
)

func TestGreeterServerImpl(t *testing.T) {
	g := grpc.NewGreeterServer()
	r, err := g.SayHello(context.Background(), &pb.HelloRequest{Name: "Alice"})

	require.Nil(t, err)
	require.Equal(t, r.Message, "Hello Alice")
}
