package main

import(
	 "fmt"
	 "math"
)

//take pointer out when using info(c) will let program run properly
//with pointer use fmt.Println(c.area())

type circle struct {
	radius float64
}

type shape interface{
	area() float64
}

func (c circle) area() float64 {   // pointer *circle
	return math.Pi * c.radius * c.radius
}
func (e circle) area1() float64 { // pointer *circle
	return math.Pi * e.radius * e.radius
}

func info(s shape) {
	fmt.Println("--------------------------")
	fmt.Println("Circle program with math.Pi")
	fmt.Println("--------------------------")
	fmt.Println(`the area is:`, s.area())
	fmt.Println("--------------------------")
	fmt.Printf("%b\t%#x\t", s.area(), s.area())
	fmt.Println("")
}


func info1(s shape) {
	fmt.Println("")
	fmt.Println(`the area is:`, s.area())
	fmt.Println("--------------------------")
	fmt.Printf("%b\t%#x\t", s.area(), s.area())
}

func main() {

       c := circle{5}
       info(c) // works without pointer
      //fmt.Println(c.area()) //works with pointer

	   e := circle{10}
	   info1(e) // works without pointer
}
