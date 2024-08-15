package main

import (
	pb "github.com/Katherine-988/subscription_proto/proto"
	"github.com/Katherine-988/subscription_server/api"
	"github.com/Katherine-988/tools"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT = ":50052"

func main() {
	// 1.创建监听器
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("init err:", err)
	}
	// 2.创建grpc server
	srv := grpc.NewServer()
	// 3.注册服务
	log.Println("正在注册服务...")
	pb.RegisterSubscriptionServiceServer(srv, &api.SubscriptionService{})

	tools.DBMgr.Init()
	tools.KafkaMgr.Init()

	// 4.启动服务
	log.Println("即将启动服务...")

	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}
