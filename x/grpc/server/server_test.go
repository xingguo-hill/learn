package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	pb "learn/x/grpc/helloworld" // 导入生成的 pb 文件

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// 从响应中提取返回的元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("无法提取元数据")
	}
	authToken := md.Get("auth-token")
	traceId := md.Get("trace-id")
	return &pb.HelloReply{
		Name: "MyName:" + in.Who + "; ctx: auth-token=" + authToken[0] + "&trace-id=" + traceId[0],
		Age:  in.YourAge}, nil
}
func (s *server) SayHelloKeyMap(ctx context.Context, in *pb.HelloRequest) (*pb.KeyValueMap, error) {
	a := make([]*pb.HelloReply, 0)
	a = append(a, &pb.HelloReply{Name: "MyName1:" + in.Who, Age: in.YourAge})
	a = append(a, &pb.HelloReply{Name: "MyName2:" + in.Who, Age: in.YourAge + 1})
	return &pb.KeyValueMap{MyMap: a}, nil
}
func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestServerTls(t *testing.T) {

	pemDir := "../tls/pem/"
	//证书配置
	certificate, err := tls.LoadX509KeyPair(pemDir+"server.crt", pemDir+"server.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(pemDir + "ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
