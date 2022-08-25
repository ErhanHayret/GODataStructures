package main

import("fmt")

type Node struct{
	Item any
	Link *Node
}

type Stack struct{
	StackHeader *Node
	Count int 
}

func newNode(item any) *Node{
	p := new(Node)
	p.Item = item
	p.Link= nil
	return p
}

func (S *Stack) Push(item any){
	newNode := newNode(item)
	if S.Count > 0 {
		newNode.Link = S.StackHeader
		S.StackHeader = newNode
	} else {
		S.StackHeader = newNode
	}
	S.Count++
}
func (S *Stack) Pop() *Node{
	tmpNode := S.StackHeader
	S.StackHeader= tmpNode.Link
	S.Count--
	return tmpNode
}

func main(){
	p:=new(Stack)
	p.Push(12)
	p.Push("as")
	p.Push("qwe")
	for p.Count!=0{
		fmt.Println(p.Count)
		fmt.Println(p.Pop())
	}
}