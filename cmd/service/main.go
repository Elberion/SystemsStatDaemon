package main

import (
	pb "SystemStatDaemon/internal/api/grpc"
	"SystemStatDaemon/internal/config"
	"SystemStatDaemon/internal/service"
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	UNIX = iota
	WIN
)

var CurrentOS = 0

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Config error: %e", err)
	}

	lsn, err := net.Listen("tcp", getPort(cfg))
	if err != nil {
		log.Fatal(err)
	}

	cfg.OSType = CurrentOS
	grpcServer := grpc.NewServer()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	pb.RegisterSystemStatServer(grpcServer, service.NewService(ctx, cfg))

	if err := grpcServer.Serve(lsn); err != nil {
		log.Fatal(err)
	}

}

func getPort(config *config.Config) (port string) {
	port = config.Port
	flag.StringVar(&port, "p", "50050", "service port")
	flag.Parse()
	if port == "" {
		log.Fatal("Port is empty")
	}
	return ":" + port
}
