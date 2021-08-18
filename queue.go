package main
import "fmt"

type Queue struct{
	items []int
}

//Enqueue

func (q *Queue) Enqueue(i int){
	q.items=append(q.items, i)
}

//Deqeue
func (q *Queue) Deqeue() int{
	to_remove:=q.items[0]
	q.items=q.items[1:]
	return to_remove
}

func main(){
	myQueue:=Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(20)
	myQueue.Enqueue(30)
	fmt.Println(myQueue.Deqeue())
	fmt.Println(myQueue)
}