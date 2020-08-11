package main

import(
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	  fmt.Println("")
	  fmt.Println("-----------------------")  //printing out CPU and Goroutine output
	  fmt.Println("CPU's:",runtime.NumCPU())
	  fmt.Println("Goroutine:", runtime.NumGoroutine())
	  fmt.Println("-----------------------")

	  counter := 0

	  const gs = 100
	  var wg sync.WaitGroup
	  wg.Add(gs)

	  for i := 0; i < gs; i++ {
	  	    go func() {               // run now with Go
	  	    	v := counter
	  	    	time.Sleep(time.Second)
	  	    	runtime.Gosched()
	  	    	v++
	  	    	counter = v
	  	    	wg.Done()
			}()

	  	    fmt.Println("Goroutine countdown:",runtime.NumGoroutine())  //count down Goroutine
	  }
	  wg.Wait()
	  fmt.Println("--------------")                               //Goroutine
	  fmt.Println("Goroutine", runtime.NumGoroutine())
	  fmt.Println("--------------")
	  fmt.Println("Count:",counter)
}
