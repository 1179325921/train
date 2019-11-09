package main

import "fmt"

func main() {
	//SliceInit()
	//TraversalSlice()
	RemoveElement()
}

func SliceInit()  {
//	从数组中初始化
	var a  = [4]int{1, 2, 3, 4}
	slice1 := a[1:3]
	fmt.Println(a, slice1)
	// 报错
	//slice2 := a[-1: 3]
	// 完全拷贝数组
	slice3 := a[:]
	fmt.Println(slice3)
	slice4 := a[:3]  // a[0:3]
	slice5 := a[1:]  // a[1: len(a)-1]
	fmt.Println(slice4, slice5)
	// 常用方式
	fmt.Println("=================")
	slice6 := []int{11, 22, 33, 44}
	fmt.Println(slice6)
	fmt.Println("Cap", cap(slice6),"Len", len(slice6))
	slice6 = append(slice6, 55)
	fmt.Println(slice6)
	fmt.Println("Cap", cap(slice6),"Len", len(slice6))
	slice6 = append(slice6, 66)
	fmt.Println(slice6)
	fmt.Println("Cap", cap(slice6),"Len", len(slice6))
//	使用make构造 切片
	fmt.Println("=================")
	slice7 := make([]int, 5, 6)
	fmt.Println(slice7)
	slice7 = append(slice7, 1)
	fmt.Println(slice7)

//	清空切片
	fmt.Println("==================")
	slice7 = slice7[0:0]
	fmt.Println(slice7)
}

func TraversalSlice()  {
	obj := []string{"a", "b", "c", "d", "e"}
	for i, v := range obj{
		fmt.Println("Index", i, "Value", v)
	}
}

func RemoveElement()  {
	obj := []string{"a", "b", "c", "d", "e"}
	//删除 下标为2 的元素
	obj = append(obj[:2], obj[3:]...)
	fmt.Println(obj)
	obj = []string{"a", "b", "c", "d", "e"}
	//删除 下标为2 后连续的2个元素
	obj = append(obj[:2], obj[2+2:]...)
	fmt.Println(obj)
}