package main

import(
	"fmt"
	"sync"
)
//WaitGroup是为了防止协程因为主线程死而无法运行
var wg sync.WaitGroup //定义 无需赋值
func main(){
	//没有打印 主死从随
	for i:=1;i<=5;i++{
		wg.Add(1) //协程开始的时候加1操作
		go func(n int){
			fmt.Println(n)
			wg.Done() //协程执行完成减1
		}(i)
	}
	//主线程一直在阻塞，什么时候wg减为0了，就停止
	wg.Wait()
}