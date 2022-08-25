package main

import("fmt")

type Node struct{
	Item any
	LLink *Node
	RLink *Node
}

type DoublyLinkedList struct{
	Header *Node
}

func newNode(item any){
	p := new(Node)
	p.Item = item
	p.LLink = nil
	p.RLink = nil
	return p
}


func (DL *DoublyLinkedList) AddLeft(item any){
	newNode := newNode(item)
	if DL.Header != nil{
		tmpNode := DL.Header
		for tmpNode

	}
}