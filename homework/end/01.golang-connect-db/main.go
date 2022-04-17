package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type PersonalInformation struct {
	Seq      int64  `json:"Seq  ,omitempty"  gorm:"column:Seq "`
	Account  int64  `json:"Account ,omitempty"  gorm:"column:Account"`
	Nickname string `json:"Nickname ,omitempty"  gorm:"column:Nickname"`
	Password string `json:"Password ,omitempty"  gorm:"column:Password"`
}

type ChatRecords struct {
	seq       int
	speaker   string
	listening string
	time      string
	record    string
}

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:learngo@tcp(127.0.0.1:3306)/learngo"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func createAccountManagement(conn *gorm.DB) error {
	resp := conn.Create(&PersonalInformation{
		Seq:      0,
		Account:  747366,
		Nickname: "小强",
		Password: "xiaoqiang",
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建账号管理表失败：", err)
		return err
	}
	fmt.Println("创建账号管理表成功")
	return nil
}

func createChattingRecords(conn *gorm.DB) error {
	resp := conn.Create(&ChatRecords{
		seq:       0,
		speaker:   "lesse",
		listening: "xiaoqiang",
		time:      "10:00",
		record:    "内容",
	})
	if err := resp.Error; err != nil {
		fmt.Println("创建聊天记录表失败：", err)
		return err
	}
	fmt.Println("创建聊天记录表成功")
	return nil
}

func main() {
	conn := connectDb()
	fmt.Println(conn)
	createAccountManagement(conn)
	createChattingRecords(conn)
}
