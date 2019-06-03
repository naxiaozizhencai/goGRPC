package main

import (
	"context"
	"fmt"
	"goGRPC/common/consul"
	"goGRPC/pb"
	"goGRPC/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
	"time"
)


func main() {
	consul.Init()

	// localIP := util.LocalIP()
	localIP := "127.0.0.1"

	target := fmt.Sprintf("%v://%v:%v/%v", "consul", localIP, 8500, "helloService")

	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)
	conn, err := grpc.DialContext(ctx, target,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)

	util.PanicIfError("fail to dial grpc server", err)
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	req := hello.HelloRequest{
		Name: "Tom Clay",
	}
	resp, err := client.SayHello(ctx, &req)
	util.PanicIfError("fail to call sayHello", err)
	log.Printf("resp:%v", resp.Reply)
}
