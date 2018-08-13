package main

import "fmt"

func array(a []int) int {
	slow := 0
	fast := 0
	for fast < len(a) -1 {
		slow ++
		fast += 2
	}
	return a[slow]
}

func main()  {
	test := []int{1,2,3,4,6,7}
	fmt.Println(array(test))
}
