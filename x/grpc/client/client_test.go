package grpctest

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "learn/x/grpc/helloworld" // 导入生成的 pb 文件
	"log"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
)

func TestDebug(t *testing.T) {

	lb := grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)
	keepAlive := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	}
	//insecure.NewCredentials() 无需TLS
	conn, err := grpc.Dial("127.0.0.1:50051", lb, grpc.WithKeepaliveParams(keepAlive), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := "Alice"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	//上下文变量传递
	md := metadata.New(map[string]string{
		"auth-token": "YOUR_AUTH_TOKEN_HERE",
		"trace-id":   "123456789",
	})

	// 将元数据附加到 gRPC 上下文中
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r1, err := c.SayHello(ctx, &pb.HelloRequest{Who: name, YourAge: 18})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s,%d\n", r1.Name, r1.Age)

	//数组测试
	r2, err := c.SayHelloKeyMap(context.Background(), &pb.HelloRequest{Who: name, YourAge: 18})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, s := range r2.MyMap {
		log.Printf("%s,%d\n", s.GetName(), s.GetAge()) //对数据进行特殊处理，如果没有就返回初始默认值
	}
}
func TestDebugTls(t *testing.T) {
	pemDir := "../tls/pem/"
	//公钥中读取和解析公钥/私钥对
	certificate, err := tls.LoadX509KeyPair(pemDir+"client.crt", pemDir+"client.key")
	if err != nil {
		log.Fatal(err)
	}
	//创建一组证书
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(pemDir + "ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	//解析证书
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "xg.com", // NOTE: this is required!
		RootCAs:      certPool,
	})

	lb := grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)
	keepAlive := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	}
	//insecure.NewCredentials() 无需TLS
	conn, err := grpc.Dial("127.0.0.1:50051", lb, grpc.WithKeepaliveParams(keepAlive), grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := "Alice"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	//上下文变量传递
	md := metadata.New(map[string]string{
		"auth-token": "YOUR_AUTH_TOKEN_HERE",
		"trace-id":   "123456789",
	})

	// 将元数据附加到 gRPC 上下文中
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	r1, err := c.SayHello(ctx, &pb.HelloRequest{Who: name, YourAge: 18})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s,%d\n", r1.Name, r1.Age)

	//数组测试
	r2, err := c.SayHelloKeyMap(context.Background(), &pb.HelloRequest{Who: name, YourAge: 18})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, s := range r2.MyMap {
		log.Printf("%s,%d\n", s.GetName(), s.GetAge()) //对数据进行特殊处理，如果没有就返回初始默认值
	}
}
