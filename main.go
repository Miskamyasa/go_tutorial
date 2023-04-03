package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {
	_, err := pl("Hello")
	if err != nil {
		panic(err)
	}
	var now = time.Now()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	_, err = pl(now.In(loc).Format("2006-01-02 15:04:05"))
	if err != nil {
		panic(err)
	}
}
