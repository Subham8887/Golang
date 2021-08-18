package Tree
 
import (
	"fmt"
)
 
type Node struct {
	val   int
	left  *Node
	right *Node
}

type Tree struct {
	head *Node
}

func (t *Node) Push(data int) {
	if t == nil {
		t=&Node{val:data,left:nil,right:nil}
	} else {
		curr:=t
		if curr.val>=data {
			if curr.left == nil {
				curr.left=&Node{val:data,left:nil,right:nil}
			} else {
				curr.left.Push(data)
			}
		} else {
			if curr.right == nil {
				curr.right = &Node{val:data,left:nil,right:nil}
			} else {
				curr.right.Push(data)
			}
		}
	}
}


func (t Node) inOrderTraversal(head *Node) []int {
	curr:=head
	arr:= []int{}
	if curr != nil {
		return arr
	}
	Stack := []*Node{}
	Stack =append(Stack, curr)
	l:=len(Stack)
	for l>0{
		if curr != nil{
			Stack=append(Stack, curr.left)
			curr=curr.left
		} else if l>0 {
			curr =Stack[l-1]
			arr = append(arr, curr.val)
			curr=curr.right
		} else {
			break
		}
		l=len(Stack)
	}
	return arr

}


func tree() *Node {
	head:=&Node{}
	return head
}


func main(){
	head:=Node{}
}