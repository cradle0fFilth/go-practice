package main

import (
	"fmt"
)

func main() {
	//默认情况下，管道是双向的--->可读可写
	//声明为只写
	var intChan2 chan<- int
	intChan2 = make(chan int, 3)
	intChan2 <- 20
	//读操作报错
	//num:= <- intChan2
	fmt.Println("intChan2", intChan2)
	//声明为只读:
	var intChan3 <-chan int
	if intChan3 != nil {
		//空管道
		num1 := <-intChan3
		fmt.Println("num1: ", num1)
	}
	//只读不可写 所以报错
	//intChan3<- 30
}
