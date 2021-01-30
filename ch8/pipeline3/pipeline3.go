package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x <= 100; x++ {
		out <- x
	}
	close(out)
}

// chan <- int 是一个只能发送的通道
// 允许发送但不允许接收

// <- chan int 是一个只能接收的int通道
// 允许接收但是不能发送
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
