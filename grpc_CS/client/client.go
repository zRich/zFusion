package main

import (
	"context"
	"fmt"
	pb "grpc_CS/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

// conversations 调用服务端的Conversations方法
func conversations() {
	//调用服务端的Conversations方法，获取流
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	for n := 0; n < 5; n++ {
		var ques string
		fmt.Println("请输入你要发送的信息：")
		fmt.Scanln(&ques)
		if ques == "exit" {
			break
		}
		err := stream.Send(&pb.StreamRequest{Question: ques})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Answer)
	}
	//最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v", err)
	}
}

// Address 连接地址
const Address string = ":8000"

var streamClient pb.StreamClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	streamClient = pb.NewStreamClient(conn)
	conversations()
}
