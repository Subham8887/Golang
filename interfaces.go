package main
import (
	"fmt"
	"math"
)

type circle struct{
	radius float64
}

type rectange struct{
	length float64
	breadth float64
}

type shape interface{
	area() float64
	peri() float64
}

func (c circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func (r rectange) area() float64{
	return r.breadth*r.length
}

func (c circle) peri() float64{
	return 2*math.Pi*c.radius
}

func (r rectange) peri() float64{
	return 2*(r.length+r.breadth)
}


func eval(g shape){
	fmt.Println(g.area())
	fmt.Println(g.peri())
}
func main(){
	circle1:=circle{
				radius: 23.3,
			}
	rectange1:=rectange{
				length: 23,
				breadth: 12,
			}
	shapes:=[]shape{circle1,rectange1}
	// area_cir:=circle1.area()
	// area_rec:=rectange1.area()
	// fmt.Println("circle area=",area_cir,"\nrectangel area=",area_rec)
	for _,shape := range shapes{
		fmt.Println(shape.peri())
	}
	// eval(rectange1)
}