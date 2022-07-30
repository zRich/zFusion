package peer

import (
	context "context"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hyperledger/fabric/common/flogging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCServer struct {
	address           string
	listener          net.Listener
	server            *grpc.Server
	lock              *sync.Mutex
	serverCertificate atomic.Value
	healthServer      *health.Server
}

var logger = flogging.MustGetLogger("PeerServer")

type PeerServer struct {
	UnimplementedEndorserServer
}

func (s *PeerServer) ProcessProposal(ctx context.Context, proposal *SignedProposal) (*ProposalResponse, error) {
	logger.Infof("ProcessProposal")
	return &ProposalResponse{
		Version:   0,
		Timestamp: &timestamppb.Timestamp{},
		Response: &Response{
			Status:  500,
			Message: fmt.Sprintf("Say hello by %s", proposal.GetProposalBytes()),
		},
		Payload: []byte(fmt.Sprintf("Your payload: %s", proposal.GetProposalBytes())),
		Endorsement: &Endorsement{
			Endorser:  []byte(fmt.Sprintf("Your payload: %s", proposal.GetProposalBytes())),
			Signature: []byte("Signed by Rich"),
		},
	}, nil
}

func NewGRPCServer(address string, serverConfig ServerConfig) (*GRPCServer, error) {
	if address == "" {
		return nil, errors.New("missing server address")
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return NewGRPCServerFromListener(lis, serverConfig)
}

func NewGRPCServerFromListener(listener net.Listener, serverConfig ServerConfig) (*GRPCServer, error) {
	grpcServer := &GRPCServer{
		address:           listener.Addr().String(),
		listener:          listener,
		server:            &grpc.Server{},
		lock:              &sync.Mutex{},
		serverCertificate: atomic.Value{},
		healthServer:      &health.Server{},
	}

	var serverOpts []grpc.ServerOption

	serverOpts = append(serverOpts, grpc.ConnectionTimeout(10*time.Second))
	grpcServer.server = grpc.NewServer(serverOpts...)

	RegisterEndorserServer(grpcServer.server, &PeerServer{})

	return grpcServer, nil
}

func (s *GRPCServer) Start() error {
	return s.server.Serve(s.listener)
}

func (s *GRPCServer) Server() *grpc.Server {
	return s.server
}

func (s *GRPCServer) Stop() {
	s.server.Stop()
}
