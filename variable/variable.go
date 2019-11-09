package main

import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	//variableBase()
	//variableQuick()
	//variableAnonymous()
	//variablePointer()
//	交换 变量的值
	a, b := 1, 2
	a, b = b, a
	fmt.Printf("a: %d, b: %d", a, b)
}

func variableBase()  {
	//	声明变量
	var s string
	var b bool
	var i int
	var u User
	var pu *User
	//	打印  未被初始化的对象，打印的未默认值, 结构体对象默认未空, 指针对象默认为 nil
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(i)
	fmt.Println(u)
	fmt.Println(pu)
}

func variableQuick()  {
	//	快捷方式 -- 最常见
	s1, b1 := "init string", true
	i1 := 123
	//  格式化输入输出 http://c.biancheng.net/view/41.html
	fmt.Printf("S1: %s, B1: %t, I1: %d \n", s1, b1, i1)
}

func variableAnonymous()  {
	//	匿名变量
	result1, _ := a()
	fmt.Printf("Result1: %d \n", result1)
}

func variablePointer()  {
	result1 := 1
	//指针变量
	//声明指针变量
	var ptrResult1 *int
	// 对result1取址并赋值给ptr_result1
	ptrResult1 = &result1
	if nil != ptrResult1 {
		//取值 *ptr_result1
		fmt.Printf("Ptr_result1 : %p, value: %d \n", ptrResult1, *ptrResult1)

	}
	//	指针的作用
	fmt.Println("ChangeValue invoke...")
	changeValueByPtr(ptrResult1)
	fmt.Printf("result1 : %d, ptr_result1: %d \n", result1, *ptrResult1)
	input1 := 1
	changeValue(input1)
	fmt.Println(input1)
}

func changeValue(in int)  {
	in = 2
	fmt.Printf("changeValue -- in: %d \n", in)
}

func changeValueByPtr(in *int)  {
	fmt.Printf("changeValueByPtr In: %d \n", *in)
	*in = 2
}

func a() (int, int)  {
	return 1, 2
}