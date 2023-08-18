package main

import (
	"fmt"
	"time"
)

func main0() {
	fmt.Println("start sub()")
	ch := make(chan string)
	go func () {
		time.Sleep(5*time.Second)
		ch <- "."
	}()
	<- ch
	fmt.Println("timeout")
}

func main () {
	<-time.After(time.Second*2)
    fmt.Println("timeout")

}