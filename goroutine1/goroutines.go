package main

import (
	"fmt"
	"time"
)

//主线程
func main() {
	
	for i:=1;i<10;i++{
		//传n每一次运行是一个新的副本
		go func(n int){
			//i是一个多协程共享数据
			fmt.Println(n)
			//匿名函数+外部变量=闭包
			//闭包
			fmt.Println(i)
		}(i)
	}
	
	time.Sleep(time.Second*2)
}
