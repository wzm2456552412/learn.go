package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn.go/pkg/apis"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewRankServiceClient(conn)
	Register(err, c)

	Login(err, c)

	GetUsers(err, c)

	defer conn.Close()
	fmt.Println("连接成功，开始聊天吧：")
	for {
		r := bufio.NewReader(os.Stdin)
		input, _, _ := r.ReadLine() // todo handle error
		if len(input) != 0 {
			talk(conn, string(input))
		}
	}
}

func Register(err error, c apis.RankServiceClient) {
	ret, err := c.ChatRegister(context.TODO(), &apis.PersonalInformation{Account: 747366})
	if err != nil {
		log.Fatal("注册失败：", err)
	}
	log.Println("注册成功:", ret)
}

func Login(err error, c apis.RankServiceClient) {
	chat, err := c.Login(context.TODO(), &apis.PersonalInformation{Account: 747366, Nickname: "小强"})
	if err != nil {
		log.Fatal("登录失败：", chat)
	}
	log.Println("登录成功，欢迎你回来，小强：", chat)
}

func GetUsers(err error, c apis.RankServiceClient) {
	list, err := c.GetOnlineUsers(context.TODO(), &apis.Null{})
	if err != nil {
		log.Fatal("获取在线用户失败：", err)
	}
	log.Println("在线用户：", list.String())
}

func talk(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Println("发送消息失败：", err)
	} else {
		data := make([]byte, 1024)
		validLen, err := conn.Read(data)
		if err != nil {
			log.Println("WARNING:读取服务器返回数据时出错：", err)
		} else {
			validLen := data[:validLen]
			log.Println("去：", message, "回：", string(validLen))
		}
	}
}
