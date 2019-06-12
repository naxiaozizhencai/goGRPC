package service

import (
	"context"
	"fmt"
	"goGRPC/pb"
	"goGRPC/util"
	"log"
	"strings"
	"time"
)

type HelloServiceImpl struct {
}


func (srv *HelloServiceImpl) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	reply := fmt.Sprintf("hello %v!", strings.ToUpper(req.Name))
	log.Printf("reply data:%v", reply)
	//mock a panic，试验server是否会挂掉
	// log.Panicf("there is a panic in sayhello")
	go util.HandlePanic(func() {
		time.Sleep(time.Millisecond * 500)
		log.Panicf("give you a panic in groutine, haha")
	})
	return &hello.HelloResponse{Reply: reply}, nil

}
