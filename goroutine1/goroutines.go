package main

import (
	"fmt"
	"time"
)

//主线程
func main() {
	//匿名函数
	for i:=1;i<10;i++{
		go func(){
			fmt.Println(i)
		}()
	}
	
	time.Sleep(time.Second*2)
}
