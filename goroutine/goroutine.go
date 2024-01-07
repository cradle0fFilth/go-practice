package main

import (
	"fmt"
	"net/http"
	"time"
)

func getFromBaidu() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Status)
}

func add(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("y can not be zero")
	}
	return x / y, nil
}

func main() {
	go func() {
		fmt.Println("Hello from a goroutine")
	}()

	go getFromBaidu()
	go http.ListenAndServe(":8080", nil)
	time.Sleep(time.Second)
}
