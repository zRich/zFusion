package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/zRich/zFusion/peer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	peer.UnimplementedEndorserServer
}

var (
	port = flag.Int("port", 3600, "The server port")
	addr = flag.String("addr", "localhost:3600", "the address to connect to")
)

func (s *server) ProcessProposal(ctx context.Context, proposal *peer.SignedProposal) (*peer.ProposalResponse, error) {
	// b := proposal.GetProposalBytes()
	// return &ProposalResponse{
	// 	Version: 1,
	// 	Response: &Response{
	// 		Payload: b,
	// 	},
	// 	Endorsement: &Endorsement{
	// 		Signature: []byte(" by Rich"),
	// 	},
	// }, nil

	return &peer.ProposalResponse{
		Version:     0,
		Timestamp:   &timestamppb.Timestamp{},
		Response:    &peer.Response{Status: 500, Message: fmt.Sprintf("Say hello by %s", proposal.GetProposalBytes())},
		Payload:     []byte(fmt.Sprintf("Your payload: %s", proposal.GetProposalBytes())),
		Endorsement: &peer.Endorsement{Endorser: []byte(fmt.Sprintf("Your payload: %s", proposal.GetProposalBytes())), Signature: []byte("Signed by Rich")},
	}, nil
}

func main() {

	// flag.Parse()
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// peer.RegisterEndorserServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	cfg := &peer.ServerConfig{}

	s, err := peer.NewGRPCServer(fmt.Sprintf(":%d", 3600), *cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}
	s.Start()
}
