package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/zRich/zFusion/peer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:3600", "the address to connect to")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := peer.NewEndorserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.ProcessProposal(ctx,
		&peer.SignedProposal{
			ProposalBytes: []byte("It's raining."),
			Signature:     []byte("Rich at Client"),
		})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Mssage : %s, Signature: %s",
		r.Endorsement.GetEndorser(),
		r.Endorsement.GetSignature())
}
