package main

import (
	"fmt"
)

func main() {
	//定义管道、声明管道------>定义一个int类型的管道
	var intChan chan int
	//通过make初始化:管道可以存放3个int类型的数据(管道存放资源)
	intChan = make(chan int, 3)
	//证明管道是引用类型
	fmt.Printf("intChain的值:%v", intChan)
	//向管道存放数据:
	intChan<- 10
	num:=20
	intChan<- num
	intChan<- 40

	//输出管道的长度：
	fmt.Printf("管道的实际长度：%v,管道的容量是:%v",len(intChan),cap(intChan))
}
