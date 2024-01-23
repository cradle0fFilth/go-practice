package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func closure1() {
	for _, salution := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salution)
		}()
	}

}

func closure2(){
	for _,salution:=range[]string{"hello","greetings","good day"}{
		wg.Add(1)
		go func(salution string){
			defer wg.Done()
			fmt.Println(salution)
		}(salution)
	}
}

func main() {
	//closure1()
	closure2()
	wg.Wait()
}
