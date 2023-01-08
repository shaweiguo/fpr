package server

import (
	pb "hyjj/pkg/proto"
	"sync"

	"google.golang.org/grpc"
)

type API struct {
	pb.UnimplementedQOTDServer
	addr       string
	quotes     map[string][]string
	mu         sync.Mutex
	grpcServer *grpc.Server
}

func New(addr string) (*API, error) {
	var opts []grpc.ServerOption
	a := &API{
		addr:   addr,
		quotes: map[string][]string{
			// TODO inset quotes here
		},
		grpcServer: grpc.NewServer(opts...),
	}
	a.grpcServer.RegisterService(&pb.QOTD_ServiceDesc, a)
	return a, nil
}
