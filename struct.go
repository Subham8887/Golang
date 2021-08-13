package main

import ("fmt")

type person struct{
	first string
	last string
	age int
}

type Gender struct{
	person
	gender bool
}

func main(){
	p1:=person{
		first: "subham",
		last: "Kumar",
		age: 22,
	}

	g1:= Gender{
		person: person{
			first: "akash",
			last: "kumar",
			age: 23,
		},
		gender: true,

	}

	fmt.Println(p1.first,p1.last,p1.age)

	fmt.Println(g1.person.first,g1.person.last)
}