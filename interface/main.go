package main

import "fmt"

type i interface {
	less(b i) bool
}

type than struct {
	a int
}

func(t *than) less(b i) bool {
	return t.a < b.(*than).a
}

type node struct {
	data interface{}
	next *node
}

func main() {
	v1 := than {
		1,
	}
	v2 := than{
		2,
	}
	fmt.Println(v1.less(&v2))

	a := &node{
		1,
		nil,
	}
	fmt.Println(a)
	b := &node{
		2,
		nil,
	}
	b.next = a
	a = b
	fmt.Println(a.next, a)
}