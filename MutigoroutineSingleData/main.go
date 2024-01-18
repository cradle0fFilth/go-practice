package main

import(
	"fmt"
	"sync"
)
var totalNum int
var wg sync.WaitGroup
var lock sync.Mutex
func add(){
	defer wg.Done()
	for i:=0;i<10000;i++{
		//加锁
		lock.Lock()
		totalNum = totalNum+1
		//解锁
		lock.Unlock()
	}
}

func sub(){
	defer wg.Done()
	for i:=0;i<10000;i++{
		totalNum = totalNum-1
	}

}

func main(){
	wg.Add(2)
	go add()
	go sub()
	//阻塞整个主线程
	wg.Wait()
	//理论上结果为0
	fmt.Println(totalNum)
}