package e2e

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	appGrpc "github.com/hamakn/go_grpc_sample/src/interfaces/grpc"
	pb "github.com/hamakn/go_grpc_sample/src/proto"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

const address = ":10000"

var c pb.GreeterClient

func TestMain(m *testing.M) {
	s, err := appGrpc.NewServer(&appGrpc.Config{Address: address})
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		s.Serve()
	}()
	defer s.Shutdown()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c = pb.NewGreeterClient(conn)

	os.Exit(m.Run())
}

func TestGRPCServer(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Bob"})
	require.Nil(t, err)

	require.Equal(t, r.Message, "Hello Bob")
}
