package main

import (
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup
//写：
func writeData(intChan chan int){
	defer wg.Done()
	for i:=1;i<=50;i++{
		intChan<- i
		fmt.Println("写入的数据为: ",i)
		time.Sleep(time.Second)
	}
	//写完就关闭
	close(intChan)
}
//读：
func readData(intChan chan int){
	defer wg.Done()
	//遍历：
	for v:=range intChan{
		fmt.Println("读取数据为：",v)
		time.Sleep(time.Second)
	}
}

func main() { //主线程
	//定义管道、声明管道------>定义一个int类型的管道
	//var intChan chan int
	//通过make初始化:管道可以存放3个int类型的数据(管道存放资源)
	wg.Add(2)
	intChan := make(chan int, 50)
	//开启读和写的协程
	go writeData(intChan)
	go readData(intChan)
	//等到两个协程执行完主线程才能四
	wg.Wait()
}
