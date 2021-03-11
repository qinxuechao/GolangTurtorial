package main

import (
	"fmt"
	"time"
)

func main() {
	// this function returns the present time
	current_time := time.Now()
	fmt.Println(current_time)
	// individual elements of time can
	// also be called to print accordingly
	fmt.Printf("%d\n",current_time.Second())

	a := fmt.Sprintf("%d-%02d-%02d-%02d-%02d-%d",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second())
	fmt.Print(a)
}
