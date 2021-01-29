package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 创建中止通道
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取单个字节
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		// 什么操作也不执行
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	launch()
}

func launch() {
	fmt.Println("Rocket launch!")
}
