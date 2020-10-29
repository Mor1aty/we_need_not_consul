package main

import (
	"context"
	"fmt"
	"github.com/Mor1aty/we_need_not_consul/proto/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/memory"
	"log"
)

func main() {

	reg := memory.NewRegistry()
	server := micro.NewService(
		micro.Name("A"),
		micro.Address(":8002"),
		micro.Version("1.0"),
		micro.Registry(reg),
	)

	//r := gin.Default()
	//r.GET("/findA", func(c *gin.Context) {
	//	service, err := reg.GetService("A")
	//	if err != nil {
	//		fmt.Println(err)
	//		c.String(200, "success")
	//		return
	//	}
	//	for _, s := range service {
	//		fmt.Printf("%#v\n", s)
	//		for _, n := range s.Nodes {
	//			fmt.Printf("%#v\n", n)
	//		}
	//	}
	//	c.String(200, "success")
	//})
	//
	//webServer := web.NewService(
	//	web.Address(":8080"),
	//	web.Name("A"),
	//	web.Handler(r),
	//	web.MicroService(server),
	//)

	shh := new(SayHelloHandler)
	if err := pb.RegisterSayHelloHandlerHandler(server.Server(), shh); err != nil {
		log.Fatalf("A server register SayHello failed, err: %v", err)
	}

	server.Init()

	if err := server.Run(); err != nil {
		log.Fatalf("A server run failed, err: %v", err)
	}
}

type SayHelloHandler struct {
}

func (shh *SayHelloHandler) SayHello(ctx context.Context, req *pb.HelloReq, resp *pb.HelloResp) error {
	fmt.Println(req.Id)
	resp.Msg = "hello world"
	return nil
}
