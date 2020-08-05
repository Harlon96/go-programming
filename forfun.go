package main

import(
	"fmt"
)

type person struct {
	  first string
	  last string
	  age int
}
func main() {

	num := []int{1, 2, 3, 4, 5, 6, 7}

	s := (`Don't forget to run two -bench's and two test's with this program.`)

	i := []int{0, 1, 2, 3, 5, 8, 13, 21}

	p1 := person{
		`Harlon`,
		`Hutchison`,
		43,
	}

	fmt.Println(s)
	fmt.Println("-----------------------------")
	fmt.Println(`Fibonacci sequence:`, i)
	fmt.Println("-----------------------------")
	fmt.Println(p1)
	fmt.Println("-----------------------------")
	fmt.Println("Number sequence:", num)

}


