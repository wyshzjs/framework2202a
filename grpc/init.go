package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RegisterGrpc(port int, server func(s grpc.ServiceRegistrar)) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	server(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
