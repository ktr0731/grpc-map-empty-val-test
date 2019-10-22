package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/ktr0731/grpc-map-empty-val-test/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var n = flag.Bool("nil", false, "return nil value with a key")

func init() {
	flag.Parse()
}

func main() {
	s := grpc.NewServer()
	api.RegisterServiceServer(s, &server{})
	reflection.Register(s)
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer s.GracefulStop()
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

type server struct{}

func (s *server) GetMap(context.Context, *api.GetMapRequest) (*api.GetMapResponse, error) {
	m := map[string]*api.Message{}
	if *n {
		m["foo"] = nil
	} else {
		m["foo"] = &api.Message{}
	}
	return &api.GetMapResponse{M: m}, nil
}
