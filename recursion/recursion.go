package main

import "fmt"

func generate() func() int {
	v1, v2 := 0,1
	return func() int {
		v1, v2 = v2, v1 + v2
		fmt.Printf("address v1 = %v, v2 = %v \n", &v1, &v2)
		return v1
	}
}

func main() {
	f := generate()
	for i := 0; i < 15; i ++ {
		fmt.Printf("the address of f is: %v", f)
		fmt.Printf("f() is %d \n", f())
	}
}
