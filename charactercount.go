package main

import(
	"fmt"
)
func main() {

	// this program is to print out the ASCII characters from 32 to 90...

	   xi:= (`Having fun with the Go language! `)

	   	  for i := 32; i <= 90; i++ {       // for init;condition;post{}
	   	  	fmt.Println("----------------------------")
	   	  	fmt.Printf("%#v\t%#x\t%U\n", i,i,i)
		  }
	   fmt.Println("----------------------------")
	   fmt.Println(xi)
	   fmt.Println("----------------------------")
}
