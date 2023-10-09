package grpctest

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	pb "learn/x/grpc" // 导入生成的 pb 文件

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func TestDebug(t *testing.T) {

	lb := grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`)
	keepAlive := keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             20 * time.Second,
		PermitWithoutStream: true,
	}
	//ginsecure.NewCredentials() 无需TLS
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

	r1, err := c.SayHello(context.Background(), &pb.HelloRequest{Who: name, YourAge: 18})
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
