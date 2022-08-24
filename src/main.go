package main

import(
	"fmt"
	//LL "datastructures/packages/LinkedLists"
	S "datastructures/packages/Stack"
)

func main(){
	//LL
	// p := new(LL.LinkedList)
	// p.AddNode("var")
	// p.PrintAllList()

	//STACK
	p:=new(S.Stack)
	p.Push(12)
	p.Push("as")
	p.Push("qwe")
	for p.Count!=0{
		fmt.Println(p.Count)
		fmt.Println(p.Pop())
	}
}