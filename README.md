# goGRPC

## 1. 功能实现
+ 实验性的Hello服务
+ 集成consul:健康检查，服务注册
+ 利用grpc负载均衡器实现client服务发现

## 2. 如何使用

### 1. build proto file

> protoc -I proto/ proto/hello.proto --go_out=plugins=grpc:pb

### 2. dev模式运行consul

> consul agent -dev

### 3. 运行server

> make run_server

### 4. 运行client

> make run_client