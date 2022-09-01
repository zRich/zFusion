package main

// server

import (
	"fmt"
	"net"
	"zFusion/test/chat_test/proto/pb/proto_demo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// gRPC服务地址
	Address = "127.0.0.1:9988"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) ProcessRequest(ctx context.Context, in *proto_demo.SignedRequest) (*proto_demo.SignedResponse, error) {
	resp := new(proto_demo.SignedResponse)
	requestBytes := "This is RequestBytes from server."
	signature := "This is Signature from server."
	resp.RequestBytes, resp.Signature = []byte(requestBytes), []byte(signature)
	return resp, nil
}

func main() {
	fmt.Println("Do you want to send REQUEST? (y/n)")
	var ans string
	fmt.Scan(&ans)
	if ans == "y" || ans == "Y" {
		conn, err := grpc.Dial(Address, grpc.WithInsecure())
		if err != nil {
			grpclog.Fatalln(err)
		}
		defer conn.Close()

		c := proto_demo.NewProcessClient(conn)

		req := &proto_demo.SignedRequest{
			RequestBytes: []byte("This is RequestBytes from client"),
			Signature:    []byte("This is Signature"),
		}
		res, err := c.ProcessRequest(context.Background(), req)
		if err != nil {
			grpclog.Fatalln(err)
		}

		fmt.Println(string(res.RequestBytes))
		fmt.Println(string(res.Signature))
	} else {
		listen, err := net.Listen("tcp", Address)
		if err != nil {
			grpclog.Fatalf("Failed to listen: %v", err)
		}

		s := grpc.NewServer()

		proto_demo.RegisterProcessServer(s, HelloService)
		fmt.Println("Listen on " + Address)
		grpclog.Println("Listen on " + Address)
		s.Serve(listen)
	}
}
