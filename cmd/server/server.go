package main

import (
	"flag"
	"fmt"
	"goGRPC/common/consul"
	"goGRPC/model"
	"goGRPC/pb"
	"goGRPC/service"
	"goGRPC/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

var (
	port = flag.Int("p", 8090, "grpc port")
)

func init() {
	flag.Parse()
	flag.VisitAll(func(i *flag.Flag) {
		log.Printf("%v:%v", i.Name, i.Value)
	})
}

func main() {

	server := grpc.NewServer()

	helloService := &service.HelloServiceImpl{}
	healthService := &service.HealthImpl{}

	hello.RegisterHelloServiceServer(server, helloService)
	grpc_health_v1.RegisterHealthServer(server, healthService)

	//注册服务到consul
	consulRegister := &model.ConsulRegister{}
	err := util.LoadJSON("conf/register.json", consulRegister)
	util.PanicIfError("fail to load consul register conf", err)
	log.Printf("register server conf:%v", consulRegister.Register)
	register := consul.NewConsulRegister(consulRegister.Register)
	//实际端口可能和配置文件不一致，所以这里赋值一下
	register.Port = *port
	if err := register.Register(); err != nil {
		util.PanicIfError("fail to register service to consul", err)
	}

	sock, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	util.PanicIfError("fail to listen port", err)

	log.Printf("grpc server start...")

	if err := server.Serve(sock); err != nil {
		util.PanicIfError("fail to start grpc server", err)
	}
}
