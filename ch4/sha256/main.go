package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	ch1 := sha256.Sum256([]byte("x"))
	ch2 := sha256.Sum256([]byte("X"))
	// %t 输出一个布尔值
	// %T 输出一个值的类型
	fmt.Printf("%x\n%x\n%t\n%T\n", &ch1, &ch2, ch1==ch2, &ch1)
	//fmt.Print()
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}