package main

import (
	"fmt"
)

func main() {
	p0 := 1
	p1 := 2
	p3 := 3
	numset1 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	numset2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("number 1:", p0, "times 5:", p0*5)
	fmt.Println("number 2:", p1, "times 5:", p1*5)
	fmt.Println("number 3:", p3, "times 5:", p3*5)

	if p0+1 == p1 {
		fmt.Println("")
		fmt.Println("both of the answers are equal to two...", p0+1, p1)
	} else {
		fmt.Println("")
		fmt.Println("answers are not the same..", p0+1, p1)
	}

	if p1 != p3 {
		fmt.Println("")
		fmt.Println("these two do not match:", p1, p3)
	} else {
		fmt.Println("")
		fmt.Println("these two numbers match:", p1, p3)
	}

	if numset1 != nil {
		fmt.Println("")
		fmt.Println("number sets are here next to each other:", numset1, numset2)
	} else {
		fmt.Println("")
		fmt.Println("number sets do not equal to each other:", numset2, numset1)
	}

	if numset2 != nil {
		fmt.Println("")
		fmt.Println("number sets are here next to each other:", numset2, numset1)
	} else {
		fmt.Println("")
		fmt.Println("numbers sets do not equal to each other:", numset1, numset2)
	}

	numset1 = numset2 //conversion

	for i, v := range numset1 {
		fmt.Println("")
		fmt.Println("index:", i, "value:", v)
	}

	switch {
	case (43 == 43):
		fmt.Println("")
		fmt.Println("data match:", numset1, numset2)
	default:
		fmt.Println("")
		fmt.Println("data does not match:", numset1, numset2)
	}

	num1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num2 := []int{11, 12, 13, 23, 24, 45, 56, 78}
	num2 = append(num1)

	fmt.Println("Appended number set:", num2)
	fmt.Println("")
	fmt.Printf("Appended: %T\t%#x\t", num2, num2)
}
