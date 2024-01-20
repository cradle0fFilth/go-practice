package main

import (
	"fmt"
)

func main() {
	//定义管道、声明管道
	var intChan chan int
	//通过make初始化，管道可以存放3个int类型的数据
	intChan = make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i
	}
	//在遍历前如果没有关闭管道,就会出现deadlock的错误(总结：写完就关，关了不影响读)
	close(intChan)
	//遍历channel
	for v := range intChan {
		fmt.Println("value=", v)
	}
}
