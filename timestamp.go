package main

import "fmt"

func main() {
	x := 1
	y := 1
	fmt.Printf("%v\n", x<<8)
	fmt.Printf("%v\n", y<<16)

	fmt.Printf("%v", x<<8+y<<16)
}
