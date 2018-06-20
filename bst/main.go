package main

import (
	"fmt"
)

type Item interface {
	Less(than Item) bool
}

type Data struct {
	value int
	id string
}

func (i Data)Less (data Item) bool {
	return i.value < data.(Data).value
}

type Tree struct {
	Root *Node
}

type Node struct {
	Left *Node
	Right *Node
	Key Item
}

func NewTree(node *Node)*Tree {
	return &Tree{
		Root:node,
	}
}

func NewNode(i Item) *Node {
	return &Node{
		Key:i,
	}
}

func (t *Tree) Insert(node *Node) {
	if t.Root == node {
		return
	}
	t.Root.insert(node)
}

func (n *Node) insert(node *Node) *Node {
	if n == nil {
		return node
	}
	if n.Key.Less(node.Key) {
		n.Right = n.Right.insert(node)
	} else {
		n.Left = n.Left.insert(node)
	}
	return n
}

func (tr *Tree) String() string {
	return tr.Root.String()
}

func (nd *Node) String() string {
	if nd == nil {
		return "[]"
	}
	s := ""
	if nd.Left != nil {
		s += nd.Left.String() + " "
	}
	s += fmt.Sprintf("%v", nd.Key)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

type test struct {
	a string
}

func(t *test) String() string {
	return fmt.Sprintf("this is a test %s", t.a)
}

func main() {
	node := NewNode(Data{4, "test2"})
	tr := NewTree(node)
	tr.Insert(NewNode(Data{2,"test1"}))
	tr.Insert(NewNode(Data{3,"test3"}))
	fmt.Println(tr)
	a := &test{"zjx"}
	fmt.Println(a)
}
