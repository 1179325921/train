package main

import (
	"fmt"
	"strconv"
)

type QttUser struct {
	User       User
	QttAccount string
}

type User struct {
	Name string
	Age  int
	// 私有属性，对外不可见
	sex bool
}

func (myself *User) SayHello() {
	fmt.Println("Hello, my name is " + myself.Name + ", age is " + strconv.Itoa(myself.Age))
}

func main() {
	//StructInit()
	QttInit()
}

func StructInit() {
	user1 := User{Name: "XiaoMing", Age: 18}
	user1.SayHello()
	var user2 User
	user2.Name = "XiaoWang"
	user2.Age = 17
	user2.SayHello()
	var user3 *User = &user2
	user3.Name = "XiaoLi"
	user3.Age = 16
	user3.SayHello()
	user2.SayHello()
}

func QttInit() {
	qttUser := QttUser{User: User{Name: "QTT", Age: 80}, QttAccount: "abc"}
	qttUser.User.SayHello()
	fmt.Println(qttUser)
}
