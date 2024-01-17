package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i < 1000; i++ {
		fmt.Println("hello golang", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

//主线程
func main() {
	go test() //开启一个协程
	for i := 1; i < 10; i++ {
		fmt.Println("hello world", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
