package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	t = t.Truncate(60 * time.Second)
	startTime := t.Unix() * 1000
	fmt.Printf("The result after rounding 't' is: %v\n", t)
	fmt.Printf("The result after rounding 'startTime' is: %v\n", startTime)

}
