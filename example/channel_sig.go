package main

import (
	"fmt"
	"time"
)

// channel实现信号通知多个Goroutines

var signal chan struct{}

func test(s chan struct{}) {
	<-s
	fmt.Printf("test")
}

func test1(s chan struct{}) {
	<-s
	fmt.Printf("test1")
}

func start(s chan struct{}) {
	fmt.Println("start start")
	go test(s)
	go test1(s)
	fmt.Println("start end")
}

func main() {
	signal := make(chan struct{})
	start(signal)
	time.Sleep(1200 * time.Second)
}
