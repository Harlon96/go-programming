package main

import(
	"fmt"
)

// this is an example of a chan  received send and quit

func main() {
	   even := make(chan int)
	   odd := make(chan int)
	   quit := make(chan int)

       x := ("about to close program integers received and sent")

	   go send(even, odd, quit)   //send channel

	   go receive(even, odd, quit)  //receive channel

	   fmt.Println("")
	   fmt.Println("---------------------------------------------------")
	   fmt.Println("about to close program integers received and sent..")
	   fmt.Println("---------------------------------------------------")
	   fmt.Println("")
	   fmt.Printf("type:%T\t\thexadecimal:%#x", x,x)
}

func receive(e,o,q <- chan int) {
	for {
		select {
		case v := <- e:
			fmt.Println("this is from the even channel:", v)
		case v := <- o:
			fmt.Println("this is from the odd channel:", v)
		case v := <- q:
			fmt.Println("the task is complete... channel closed:", v)
			return
		}
	}
}

      func send(e,o,q  chan<- int) {
      	for i := 0; i < 100; i++ {
      		if i % 2 == 0 {
      			e <- i
			} else {
				o <- i
			}
		}
		  q <- 0
	  }