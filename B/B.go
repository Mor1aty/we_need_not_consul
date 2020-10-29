package main

import (
	"context"
	"fmt"
	"github.com/Mor1aty/we_need_not_consul/proto/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/memory"
	"log"
)

func main() {
	reg := memory.NewRegistry()
	if err := reg.Register(&registry.Service{
		Name:    "A",
		Version: "1.0",
		Nodes: []*registry.Node{
			{
				Address: "10.81.9.94:8002",
			},
		},
	}); err != nil {
		log.Fatalf("registry failed, err: %v", err)
	}
	server := micro.NewService(
		micro.Registry(reg),
	)

	shhs := pb.NewSayHelloHandlerService("A", server.Client())
	resp, err := shhs.SayHello(context.Background(), &pb.HelloReq{
		Id: 10,
	})

	if err != nil {
		log.Fatalf("SayHello request failed, err: %v", err)
	}
	fmt.Println(resp)
}
