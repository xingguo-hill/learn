package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"testing"

	pbhello "learn/x/grpc/helloworld" // 导入生成的 pb 文件
	pb "learn/x/grpc/helloworldgw"    // 导入生成的 pb 文件

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

type servergw struct {
	pb.UnimplementedGreeterGWServer
}

func (s *servergw) SayHelloGWJSon(ctx context.Context, in *pbhello.HelloRequest) (*pb.HelloReplyGW, error) {
	log.Printf("%v,%v", in.Who, in.YourAge)
	return &pb.HelloReplyGW{MyName: "Post MyName:" + in.Who, Age: in.YourAge}, nil
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
	//封送处理器
	opt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:  true, //使用.proto变量设置，默认为false会转成驼峰
			UseEnumNumbers: true,
		},
	})
	gwmux := runtime.NewServeMux(opt)

	//不设置头信息
	// gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = pb.RegisterGreeterGWHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	//与http.ListenAndServe(":81", gwmux)实现类似的功能，区别可以配置更多的http参数
	gwServer := &http.Server{
		Addr:    ":81",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:81")
	log.Fatalln(gwServer.ListenAndServe())
}
