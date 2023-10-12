package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"testing"

	pb "learn/x/grpc" // 导入生成的 pb 文件

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type servergw struct {
	pb.UnimplementedGreeterGWServer
}

func (s *servergw) SayHelloGWJSon(ctx context.Context, in *pb.HelloRequestGW) (*pb.HelloReplyGW, error) {
	log.Printf("%v,%v", in.Who, in.YourAge)
	return &pb.HelloReplyGW{Name: "Post MyName:" + in.Who, Age: in.YourAge}, nil
}
func (s *servergw) SayHelloGWGet(ctx context.Context, in *pb.HelloGetRequestGW) (*pb.HelloGetReplyGW, error) {
	return &pb.HelloGetReplyGW{Msg: "Get id:" + cast.ToString(in.GetId())}, nil
}
func TestGWServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterGWServer(s, &servergw{})
	//启动grpc端口
	log.Println("Serving gRPC on 0.0.0.0:50051")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = pb.RegisterGreeterGWHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":81",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:81")
	log.Fatalln(gwServer.ListenAndServe())
}
