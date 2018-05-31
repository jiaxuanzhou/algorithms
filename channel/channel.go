package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func chanA(stopC <-chan struct{}) {
	go func(stopCh <- chan struct{}) {
		<- stopCh
		fmt.Println("stop singal 1")
	}(stopC)

	go wait.Until(func(){
		fmt.Println("this is a test")
	},time.Second *2, stopC)
	<- stopC
	fmt.Println("recieved stop signal")
}

func main() {

	testC := make(chan struct{})
	hangC := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		close(testC)
	}()
	chanA(testC)
	<- hangC
}
