package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/zRich/zFusion/peer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedEndorserServer
}

var (
	port = flag.Int("port", 3600, "The server port")
	addr = flag.String("addr", "localhost:3600", "the address to connect to")
)

func main() {

	// addr := "127.0.0.1:3600"
	// cfg := &network.ServerConfig{}
	// srv, err := network.NewGRPCServer(addr, *cfg)
	// pb.RegisterEndorserServer(srv.Server(), &server{})
	// if err != nil {
	// 	go srv.Start()
	// }

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEndorserServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	go RunClient()
}

func RunClient() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEndorserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.ProcessProposal(ctx,
		&pb.SignedProposal{
			ProposalBytes: []byte("hello world"),
		})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Endorsement.GetSignature())
}
