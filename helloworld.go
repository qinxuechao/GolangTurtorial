package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1"
	b,_ := strconv.ParseFloat(a, 64)
	fmt.Println(b)
	fmt.Println(int64(b * 60 * 1000))

}


