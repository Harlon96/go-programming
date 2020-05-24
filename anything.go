package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most awesome class in the entire world for having fun " +
		"and learning the Go programming language.")
	foo()
	fmt.Print("something else and so forth")

	for i := 0; i < 20; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}
func foo() {
	fmt.Println("I like the FooFighters.")
}
func bar() {
	fmt.Println("and we exited baby!!")
}

//  control flow:
//  (1) sequence
//  (2) loop; iterative
//  (3) conditional
