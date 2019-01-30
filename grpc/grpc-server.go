package api

import (
	"google.golang.org/grpc"
)

//Server : GrpcServer struct
type Server struct {
	server *grpc.Server
}

//New : instantiate new GrpcServer
func New() (*Server, error) {
	grpcServer := grpc.NewServer()
	return &Server{server: grpcServer}, nil
}
