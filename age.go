package main

import(
	"fmt"
)

// the age program is to find out my date of birth

func main() {
	bd := 1977
    b := 1953
    a := 1977
	for {                       // for break statement
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++
		fmt.Println("----------------")
	}
	     switch bd {                              // switch case statements
		 case b:
		 	fmt.Println("False, I am forty-three years of age in 2020...")
		 case bd:
		 	fmt.Println("That is correct I am forty-three years of age this January 9th, 2020...")
		 case a:
			 fmt.Printf("That is correct, I am:%T\t%d\t%#x\t", bd,bd,bd)
		 }

}

