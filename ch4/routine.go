package main

import (
	"fmt"
	"time"
)

// 新しく作られるgoroutineが呼ぶ関数
func sub() {
	fmt.Println("sub() is running") 
	time.Sleep(time.Second) 
	fmt.Println("sub() is finished")
}
func main0() {
	fmt.Println("start sub()") 
	// goroutineを作って関数を実行 
	go sub()
	time.Sleep(2 * time.Second)
}

func main() {
	fmt.Println("start sub()")
	tasks := make(chan bool)

	go func() {
		fmt.Println("sub() is running")
		time.Sleep(10*time.Second)
		tasks <- true
		fmt.Println("sub() is finished")
	}()

	<-tasks
	close(tasks)
	fmt.Println("timeout")
}