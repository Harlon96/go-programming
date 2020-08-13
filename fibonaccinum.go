package main

import(
	"fmt"
)



	// variadic functions with fibonacci numbers

	func numb(nums ...int) {
		fmt.Print(nums, " ")
		total := 0
		for _, num := range nums {
			total += num
		}

		fmt.Println("sum of the numbers:", total)

	}

	func main() {
		fmt.Println("")
		fmt.Println("Fibonacci numbers:")
		fmt.Println("")

		numb(0, 1, 1, 2)
		fmt.Println("")
		numb(2, 3, 5, 8, 13, 21, 34, 55)
		fmt.Println("")
		nums := []int{5, 8, 13, 21, 34, 55}
		numb(nums...)

	}

