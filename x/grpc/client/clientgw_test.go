package grpctest

import (
	"bytes"
	"context"
	"fmt"
	"io"
	pb "learn/x/grpc"
	"log"
	"net/http"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func TestPostGRPCGW(t *testing.T) {
	// 创建 HTTP 客户端
	client := &http.Client{}

	// 准备 POST 请求的数据
	data := []byte(`{"who": "xiaoming", "your_age": 17}`) // 替换为您的请求数据

	// 创建一个请求
	req, err := http.NewRequest("POST", "http://127.0.0.1:81/api/sayHello", bytes.NewBuffer(data)) // 替换为您的目标 URL
	if err != nil {
		fmt.Println("创建请求出错:", err)
		return
	}

	// 设置请求头，如果有需要的话
	// req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP请求出错:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应出错:", err)
		return
	}

	// 输出响应内容
	fmt.Println("响应内容:")
	fmt.Println(string(body))
}
func TestGetGRPCGW(t *testing.T) {
	// 发送 GET 请求
	resp, err := http.Get("http://127.0.0.1:81/api/sayHelloGet/123")
	if err != nil {
		fmt.Println("HTTP请求出错:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应出错:", err)
		return
	}

	// 输出响应内容
	fmt.Println("响应内容:")
	fmt.Println(string(body))
}

func TestGRPC(t *testing.T) {
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
	c := pb.NewGreeterGWClient(conn)

	name := "Alice"
	r1, err := c.SayHelloGWJSon(context.Background(), &pb.HelloRequestGW{Who: name, YourAge: 18})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s,%d\n", r1.MyName, r1.Age)
}
