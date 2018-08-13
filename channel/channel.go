package main

import (
	"fmt"
	"time"
	"reflect"
	"unsafe"

	"k8s.io/apimachinery/pkg/util/wait"
)

func isChanClosed(ch interface{}) bool {
	if reflect.TypeOf(ch).Kind() != reflect.Chan {
		panic("only channels!")
	}

	// get interface value pointer, from cgo_export
	// typedef struct { void *t; void *v; } GoInterface;
	// then get channel real pointer
	cptr := *(*uintptr)(unsafe.Pointer(
		unsafe.Pointer(uintptr(unsafe.Pointer(&ch)) + unsafe.Sizeof(uint(0))),
	))

	// this function will return true if chan.closed > 0
	// see hchan on https://github.com/golang/go/blob/master/src/runtime/chan.go
	// type hchan struct {
	// qcount   uint           // total data in the queue
	// dataqsiz uint           // size of the circular queue
	// buf      unsafe.Pointer // points to an array of dataqsiz elements
	// elemsize uint16
	// closed   uint32
	// **

	cptr += unsafe.Sizeof(uint(0))*2
	cptr += unsafe.Sizeof(unsafe.Pointer(uintptr(0)))
	cptr += unsafe.Sizeof(uint16(0))
	return *(*uint32)(unsafe.Pointer(cptr)) > 0
}

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

func testC() {

	testC := make(chan struct{})
	hangC := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 10)
		close(testC)
	}()
	chanA(testC)
	<- hangC
}

func testChanClosed(ch chan struct{}) bool {
	fmt.Println("test")
	ch <- struct{}{}
	if len(ch) == 0 {
		select {
		case _, ok := <-ch:
			return !ok
		}
	}
	return false
}

func IsClosed(ch <-chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func main() {
	c := make(chan struct{})
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
}
