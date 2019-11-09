package main

import "fmt"

type People interface {
	Sleep()
	Eat() bool
	Coding(code string) error
}

type Hello interface {
	SayHello()
}

type QttProgrammer struct {
	Name string
}

func (myself *QttProgrammer) Sleep()  {
	fmt.Println(myself.Name, " zzZ")
}

func (myself *QttProgrammer)Eat() bool {
	fmt.Println(myself.Name, " eating ....")
	return true
}

func (myself *QttProgrammer)Coding(code string) error  {
	fmt.Println(myself.Name, " coding: ", code)
	return nil
}

func (myself *QttProgrammer) SayHello()  {
	fmt.Println("Hello, ", myself.Name)
}

func main() {
	qttProgrammer := &QttProgrammer{Name:"ZhaoSi"}
	OneDay(qttProgrammer)
	Say(qttProgrammer)
}

func OneDay(people People) {
	people.Eat()
	people.Coding("Hello world!")
	people.Sleep()
}

func Say(hello Hello)  {
	hello.SayHello()
}