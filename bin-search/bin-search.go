package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 6, 7}
	fmt.Println(binSearch(a, 1))
	fmt.Println(recurBinSearch(a, 0, 5, 7))
}

// return -1 if there is no element match
func binSearch(a []int, target int) int {
	low := 0
	high := len(a) - 1
	for {
		mid := low + (high-low)/2
		fmt.Printf("%d \n", mid)
		switch {
		case a[mid] > target:
			high = mid - 1
		case a[mid] < target:
			low = mid + 1
		case a[mid] == target:
			return mid
		}
		if low > high {
			break
		}
	}
	return -1
}

func recurBinSearch(a []int, low, high, target int) int {

	if low > high {
		return -1
	}
	mid := low + (high-low)/2

	if a[mid] == target {
		return mid
	}

	if a[mid] > target {
		return recurBinSearch(a, low, mid-1, target)
	}

	if a[mid] < target {
		return recurBinSearch(a, mid+1, high, target)
	}

	return -1
}
