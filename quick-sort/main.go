package main

import (
	"fmt"
)

func quickSort(a []int, length int) {
	if length < 2 {
		return
	}
	max := 0
	for i := 0; i < length; i++ {
		if a[max] <= a[i] {
			max = i
		}
	}
	a[length-1], a[max] = a[max], a[length-1]
	quickSort(a, length-1)
}

func main() {
	a := []int{2, 9, 9, 1}
	quickSort(a, 4)
	fmt.Println(a)
}
