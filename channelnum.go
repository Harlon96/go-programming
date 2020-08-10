package main

import(
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)

	for v := range c {
		  fmt.Println(v)
	}
	fmt.Println(`I have sent 25 numbers to channel c`)
	fmt.Println(`Thank you for playing...`)
}

//send channel

func send(c chan<- int) {
	for i := 0; i < 25; i++ {
		   fmt.Println(i)
		   }
		   close(c)
}