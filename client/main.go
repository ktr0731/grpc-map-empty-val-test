package main

import (
	"context"
	"log"

	"github.com/k0kubun/pp"
	"github.com/ktr0731/grpc-map-empty-val-test/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := api.NewServiceClient(conn)
	res, err := c.GetMap(context.Background(), &api.GetMapRequest{})
	if err != nil {
		log.Fatal(err)
	}
	pp.Printf("%v\n", res)
}
