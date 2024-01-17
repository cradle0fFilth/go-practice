package main 
import(
	"fmt"
	"time"
	"strconv"
)
func test(){
 for i:=1;i<10;i++{
	fmt.Println("hello golang",strconv.Itoa(i))
	time.Sleep(time.Second)
 }
}
func main(){
	go test()
	for i:=1;i<10;i++{
		fmt.Println("hello world",strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}