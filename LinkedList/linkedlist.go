package main
import "fmt"

type Node struct {
	val int
	next *Node
}

type Linkedlist struct {
	head *Node
}

func (l *Linkedlist) insert(data int ) {
	n:=Node{}
	n.val=data
	if l.head == nil{
		l.head = &n
		return
	}
	curr:=l.head
	for curr.next != nil{
		curr=curr.next
	}
	curr.next=&n
}

func (l *Linkedlist) display(){
	curr:=l.head
	for curr != nil {
		fmt.Printf("%d --> ",curr.val)
		curr=curr.next
	}
	fmt.Println()
}

func (l *Linkedlist) delete_num_all(val int){
	curr:=l.head
	
}

func main(){
	LL:=Linkedlist{}
	LL.insert(10)
	LL.insert(20)
	LL.insert(30)
	LL.display()
}

