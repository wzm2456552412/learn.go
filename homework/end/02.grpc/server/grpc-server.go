package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"learn.go/pkg/apis"
	"log"
	"net"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	startGRPCServer(ctx)
	Chat()
}

func startGRPCServer(ctx context.Context) {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer([]grpc.ServerOption{}...)
	apis.RegisterRankServiceServer(s, &rankServer{
		persons: map[string]*apis.PersonalInformation{},
	})
	go func() {
		select {
		case <-ctx.Done():
			s.Stop()
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var qa = map[string]string{
	"你好": "你好",
	"晚上要不要一起去研究多线程编程?": "好的，我也正想找你一起研究研究呢。",
}

func Chat() {
	var port string
	flag.StringVar(&port, "port", "8080", "配置启动端口")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("warning:建立连接失败：", err)
			continue
		}
		fmt.Println(conn)

		// talk(conn)
		go talk(conn)
	}
}

func talk(conn net.Conn) {
	defer fmt.Println("结束链接：", conn)
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		valid, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("WARNING:读数据失败：", err)
			continue
		}
		content := buf[:valid]
		resp, ok := qa[string(content)]
		if !ok {
			log.Println("没有找到回答，问他说了什么")
			conn.Write([]byte(`我听不懂你在说什么`)) // handle eror
			continue
		}
		conn.Write([]byte(resp)) // handle eror
		if string(content) == "再见" {
			break
		}
	}
}
