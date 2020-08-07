package main

import(
	"fmt"
)
func main() {
	c := make(chan int, 1) // #1 makes this a buffered channel

	x := `hello, my name is harlon. yours is?`
	y := `hi my name is morpheus.`
	x1 := 43
	y1 := 50

	// basic communication and age in binary

	fmt.Println(``)
	fmt.Println(x)
	fmt.Println(``)
	fmt.Println(y)
	fmt.Println(``)
	fmt.Printf("%b\n", y1)
	fmt.Println(``)
	fmt.Printf("%b", x1)
	fmt.Println(``)

	//anonymous func

	go func() {
		c <- 42
	}()
	fmt.Println(``)
	fmt.Println(<-c)
	fmt.Println(``)
	fmt.Println(c)
	fmt.Println(``)
	fmt.Printf("%T\t%b\t%d\t", c, c, c)

}

