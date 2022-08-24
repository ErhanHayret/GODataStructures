package LinkedLists

import(
	"fmt"
)

type Node struct{
	Item any
	Link *Node
}

type LinkedList struct{
	ListHead *Node
	ListTail *Node
}

func newNode(item any) *Node{
	p := new(Node)
	p.Item = item
	p.Link = nil
	return p
}

func (LL *LinkedList) AddNode(item any){
	newNode := newNode(item)
	if LL.ListHead == nil{
		LL.ListHead = newNode
		LL.ListTail = newNode
	} else {
		LL.ListTail.Link = newNode
		LL.ListTail = newNode
	}
}

func (LL *LinkedList) PrintAllList(){
	tmp := LL.ListHead
	for tmp != nil {
		fmt.Println(tmp.Item)
		tmp = tmp.Link
	}
}