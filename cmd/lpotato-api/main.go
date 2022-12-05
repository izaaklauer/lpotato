package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/izaaklauer/lpotato/config"
	lpotatov1 "github.com/izaaklauer/lpotato/gen/proto/go/lpotato/v1"
	"github.com/izaaklauer/lpotato/server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("starting lpotato.......")
	defer fmt.Println("lpotato exiting!")

	err := serve()
	if err != nil {
		log.Fatalf("failed serving\n%s", err)
	}
}

func serve() error {
	var c config.Config

	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		var err error
		c, err = config.GetConfig(configPath)
		if err != nil {
			return errors.Wrapf(err, "failed to load config from %q", configPath)
		}
	} else {
		// Erroring on no config would be totally valid
		// return errors.New("Environment variable CONFIG_PATH is unset")

		c = config.DefaultConfig()
	}

	// Start the service
	lpotatoServer, err := server.NewLpotatoServer(c.Lpotato)
	if err != nil {
		return errors.Wrapf(err, "failed to start lpotato server")
	}

	listener, err := net.Listen("tcp", c.Server.BindAddr)
	if err != nil {
		return errors.Wrapf(err, "failed to listen on %s", c.Server.BindAddr)
	}
	grpcServer := grpc.NewServer()

	lpotatov1.RegisterLpotatoServiceServer(grpcServer, lpotatoServer)
	reflection.Register(grpcServer)

	log.Printf("Serving on %q", c.Server.BindAddr)
	if err := grpcServer.Serve(listener); err != nil {
		return errors.Wrapf(err, "gRPC server exited")
	}

	return nil
}
