package main

import (
	"fmt"
)

func main() {
	MapInit()
	TraversalMap()
	RemoveElementForMap()
}

func MapInit()  {
	//Make 初始化
	map1 := make(map[int]string)
	fmt.Println(map1)
	map1[0] = "Zero"
	map1[2] = "Two"
	map1[3] = "There"
	fmt.Println(map1)
	// 直接初始化
	fmt.Println("================")
	map2 := map[string]int{"One": 1, "Two": 2, "There" : 3}
	fmt.Println(map2)
}

func TraversalMap()  {
	myMap := map[string]int{"One": 1, "Two": 2, "There" : 3}
	for k, v := range myMap {
		fmt.Println("Key", k, "Value", v)
	}
}

func RemoveElementForMap()  {
	myMap := map[string]int{"One": 1, "Two": 2, "There" : 3}
	fmt.Println(myMap)
	delete(myMap, "Two")
	fmt.Println(myMap)
}