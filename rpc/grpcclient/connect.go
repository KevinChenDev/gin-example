package grpcclient

import (
	"os"

	grpc "google.golang.org/grpc"
	pb "micros/rpc/proto"
)

var (
	client pb.MicrosServiceClient
	conn   *grpc.ClientConn
)

func Connect() pb.MicrosServiceClient {
	if client != nil {
		return client
	}
	address := os.Getenv("MICROS_GRPC_LISTEN")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = pb.NewMicrosServiceClient(conn)
	return client
}

func Close() error {
	if conn != nil {
		defer func() { client = nil }()
		return conn.Close()
	}
	return nil
}
