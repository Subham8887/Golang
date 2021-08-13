package main
import ("fmt")
type name struct{
	first string
	last string
}
type fathers_name struct{
	first string
	last string
}

type student struct{
	id int
	name
	fathers_name
	age int

}

func (s student) stud() {
	fmt.Println("Your name is",s.name.first,s.name.last,"and you age is",s.age)
	fmt.Println("Your fathers name is",s.fathers_name.first,s.fathers_name.last)
}
func main(){
	stud1:=student{
		id: 01,
		name: name{
					first: "subham",
					last: "kumar",
				},
		fathers_name: fathers_name{
					first:"sant",
					last:"prasad",
				},
		age: 22,
	}
	fmt.Println(stud1)
	stud1.stud()
}