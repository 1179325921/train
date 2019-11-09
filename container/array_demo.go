package main

import "fmt"

func main() {
	ArrayBase()
	//ArrayCompare()
}

func ArrayBase()  {
	var a [3]int  = [3]int{1,2,3}           // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"
}

func ArrayCompare()  {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // 编译错误：无法比较 [2]int == [3]int
}