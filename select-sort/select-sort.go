package main

import "fmt"

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func less(a []int, i, j int) bool{
	return a[i] < a[j]
}

func selectSortRecursive(a []int, start int) {
	if start > len(a) -1 {
		return
	}
	min := start
	for i := start + 1; i < len(a); i++ {
		if less(a, i, min) {
			min = i
		}
	}
	swap(a, start, min)
	start ++
	selectSortRecursive(a, start)
}

func selectSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i +1; j < len(a); j++ {
			if less(a, j, i) {
				min = j
			}
		}
		swap(a, i, min)
	}
}

func main() {
	a := []int{1,4,3,2,1}
	b := []int{1,4,3,2,1}
	selectSortRecursive(a,0)
	selectSort(b)
	fmt.Println(a)
	fmt.Println(b)
}