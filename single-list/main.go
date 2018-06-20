package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type Heap interface {
	Pop(node *Node) *Node
	Push(node *Node)
}

func New(data interface{}) *Node {
	return &Node{
		Data: data,
		Next:nil,
	}
}

func (n *Node) Less() bool{
	return n.Data.(int) < n.Next.Data.(int)
}



func (n *Node) InsertFront(node *Node) {
	if n == nil {
		n = node
	} else {
		temp := &Node{}
		temp.Next = n.Next
		temp.Data = n.Data
		n.Data = node.Data
		n.Next = temp
	}
}

func (n *Node) Delete(node *Node) {
	for  n.Next != nil {
		if n.Data == node.Data {
			n.Data = n.Next.Data
			n.Next = n.Next.Next
			return
		}
		n = n.Next
	}
}

func (n *Node) Get(node *Node) *Node{
	for  n.Next != nil {
		if n.Data == node.Data {
			return n
		}
		n = n.Next
	}
	return nil
}

func (n *Node) Pop(node *Node) *Node{
	res := n.Get(node)
	n.Delete(node)
	return res
}

func (n *Node) Push(node *Node) {
	n.InsertLast(node)
}

func(n *Node) InsertLast(node *Node) {
	cNode := n
	for cNode != nil {
		if cNode.Next == nil {
			cNode.Next = node
			break
		}
		cNode = cNode.Next
	}
}

func main() {
	n1 := New(1)
	n2 := New(2)
	n3 := New(3)
	n1.InsertFront(n2)
	n1.InsertFront(n3)
	//n4 := New(4)
	n1.Pop(n2)
	for t := n1; t != nil; t = t.Next {
		fmt.Println(t.Data)
	}
}
