package main

import(
	"fmt"
)

// making a chan sending numbers over that chan and closing that chan...
func main() {
	c := make(chan int)

	go send (c)

	for v := range c {     //for range loop
		 fmt.Println(v)
	}
	fmt.Println("")
	fmt.Println("----------------------------------------------")
	fmt.Println(`sending some numbers to your channel buddy...`)
	fmt.Println("----------------------------------------------")
}

//send channel
func send(c chan<- int) {

	for i := 0; i < 25; i++ {                   //for init;condition;post i++ {}
		 fmt.Println("")
		 fmt.Println("here they come:", i)
	}
	close(c)
}
