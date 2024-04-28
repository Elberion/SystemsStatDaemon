package main

import (
	pb "SystemStatDaemon/internal/api/grpc"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Service struct {
	pb.UnimplementedSystemStatServer
}

func main() {
	var port string
	flag.StringVar(&port, "p", "50050", "service port")
	flag.Parse()
	port = ":" + port

	lsn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSystemStatServer(grpcServer, new(Service))
	if err := grpcServer.Serve(lsn); err != nil {
		log.Fatal(err)
	}

}
