package server

import (
	"context"
	"log"

        "github.com/izaaklauer/lpotato/config"
	lpotatov1 "github.com/izaaklauer/lpotato/gen/proto/go/lpotato/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LpotatoServer struct {
	lpotatov1.UnimplementedLpotatoServiceServer

	config config.Lpotato
}

// NewLpotatoServer initializes a new server from config
func NewLpotatoServer(config config.Lpotato) (*LpotatoServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := LpotatoServer{
		config: config,
	}

	return &server, nil
}

func (s * LpotatoServer) HelloWorld(
	ctx context.Context,
	req *lpotatov1.HelloWorldRequest,
) (*lpotatov1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &lpotatov1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
