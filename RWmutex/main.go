package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var wg sync.WaitGroup

func read() {
	defer wg.Done()
	mu.RLock() //读数据锁不产生影响，同时读写才产生影响
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	mu.RUnlock()
}

func write() {
	defer wg.Done()
	mu.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 1)
	fmt.Println("修改数据成功")
	mu.Unlock()
}

func main() {
	wg.Add(5)
	//启动协程--->读多写少
	for i := 0; i < 5; i++ {
		go read()
	}
	//go write()
	wg.Wait()

}
