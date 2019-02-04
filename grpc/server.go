package api

import (
	"google.golang.org/grpc"
)

//Server : GrpcServer struct
type Server struct {
	Instance *grpc.Server
}

//New : instantiate new GrpcServer
func New() (*Server, error) {
	grpcServer := grpc.NewServer()
	return &Server{Instance: grpcServer}, nil
}
