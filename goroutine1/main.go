package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello golang", strconv.Itoa(i))
		//阻塞一秒
		time.Sleep(time.Second)
	}
}

func main() {
	test()
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world +" + strconv.Itoa(i))
		//阻塞一秒
		time.Sleep(time.Second)
	}
	
}
