package main

import (
	"fmt"
	"time"
)

func main()  {
	const longForm = time.RFC3339
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Print(loc)
	beiginTime := "2020-10-08 11:00:00"
	t, _ := time.ParseInLocation(longForm, beiginTime, loc)
	fmt.Println(t)

	t2, _ := time.Parse(longForm, beiginTime)
	fmt.Println(t2)

	var t3 time.Time
	t3.UnmarshalJSON([]byte(beiginTime))
	fmt.Println(t3)
}
