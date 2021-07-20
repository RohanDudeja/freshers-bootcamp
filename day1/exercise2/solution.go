package main

import "fmt"

type node struct{
	val string
	left *node
	right *node
}
func Preorder(head *node) {
	if head == nil {
		return
	}
	fmt.Print(head.val," ")
	Preorder(head.left)
	Preorder(head.right)
}

func Postorder(head *node) {
	if head == nil {
		return
	}
	Postorder(head.left)
	Postorder(head.right)
	fmt.Print(head.val," ")
}
func main(){
	// tree construction
	head:= &node{val: "+",left: nil,right: nil}
	a:=&node{val: "a",left: nil,right: nil}
	minus:=&node{val: "-",left: nil,right: nil}
	c:=&node{val: "c",left: nil,right: nil}
	b:=&node{val: "b",left: nil,right: nil}
	head.left=a
	head.right=minus
	minus.left=b
	minus.right=c
	//
	fmt.Print("Preorder: ")
	Preorder(head)
	fmt.Println()
	fmt.Print("Postorder: ")
	Postorder(head)
	fmt.Println()

}
