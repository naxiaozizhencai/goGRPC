package service

import (
	"context"
	"fmt"
	"goGRPC/pb"
	"strings"
)

type HelloServiceImpl struct {
}

func (srv *HelloServiceImpl) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	reply := fmt.Sprintf("hello %v!", strings.ToUpper(req.Name))
	return &hello.HelloResponse{Reply: reply}, nil

}
