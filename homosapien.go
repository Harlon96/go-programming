package main

import(
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

type homosapien interface {
	   speak()
}

func(p *person) speak() {
	fmt.Println("")
	fmt.Println("-----------------------------------------")
	fmt.Println(`   Hello homosapien, this is Harlo...`)
	fmt.Println("-----------------------------------------")
	fmt.Println("")
}

func saySomething(h homosapien) {
	h.speak()
	fmt.Println("")
	fmt.Println("-----------------------------------------")
	fmt.Println("saySomething is in progress.. working now")
	fmt.Println("-----------------------------------------")
	fmt.Println("")
}

func main() {
      p1 := person{
      	first: "harlo",
      	last: "hutch",
      	 age:  43,
	  }
	  saySomething(&p1)  //prints out saySomething

      p1.speak()       //this will print out (p *person) speak()

      fmt.Println(p1)

      p2 := ("harlohutch43")

	  fmt.Println("")
      fmt.Printf("type:%T\t\thexadecimal: %#x\t", p2, p2)
	  fmt.Println("")
}
