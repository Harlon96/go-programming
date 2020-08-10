package main

import(
	"fmt"
	"github.com/go-programming/code_samples/010-ninja-level-thirteen/01/finished/dog"
)

type canine struct {
	 name string
	 age int
}

func main() {
	fido := canine{
		name: "fido",
		age:  dog.Years(5),
	}
		fmt.Println(fido)
		fmt.Println(dog.YearsTwo(18))
	}