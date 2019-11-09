package main

import (
	"fmt"
	"time"
)

func Product(ch chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Println("Product:", i)
		ch <- i
	}
}

func Consumer(ch <-chan int) {
	for i := 0; i < 100; i++ {
		a := <-ch
		fmt.Println("Consmuer:", a)
	}
}

func main() {
	ch := make(chan int, 1)
	go Product(ch)
	go Consumer(ch)
	time.Sleep(5000000)
}