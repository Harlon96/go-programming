package main

import(
	"fmt"
)

//bit changing

func main() {
	  a := `Here I am just having fun with a raw string literal!!`

	  kb := 1024
	  mb := 1027 * kb
	  gb := 1024 * mb
	                                            //printing out a raw string literal...
	  fmt.Println(``)
	  fmt.Println(`-------------------`)
	  fmt.Println(a)
	  fmt.Println(`-------------------`)
	                                            //printing out bit change
	  fmt.Println(``)
	  fmt.Printf("%d\t\t%b\n", kb, kb)
	  fmt.Printf("%d\t\t%b\n", mb, mb)
	  fmt.Printf("%d\t\t%b\n", gb, gb)
	  fmt.Println(`-------------------`)
	  fmt.Println(``)

	  for i := 0; i < 10; i++ {     //printing out numbers from 0 to 9 with a for init;condition;post i++ {} loop
	  	   fmt.Println(i)
	  }

	fmt.Println(``)
	  
	n := "Harlo"
	                              //switch case default loop
	  switch n {
	  case "money":
	  	fmt.Println(`I have some money...`)
	  case `sunshine`:
	  	fmt.Println(`It is a sunny day today...`)
	  case `wet`:
	  	fmt.Println(`the water is wet!`)
	  case `n`:
	  	fmt.Println(`Harlo is enjoying his work today!`)
	  default:
	  	fmt.Printf(n)
	  	}

	  fmt.Println(``)

	  x := []int{0,1,2,3,5}            //printing a slice, appending a slice then printing out the new slice

	  fmt.Println(``)
	  fmt.Println(x)
	  fmt.Println(``)

	  x = append(x,8,13,21,34,55)

	  fmt.Println(x)
	  fmt.Println(``)

}