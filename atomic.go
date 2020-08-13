package main

import(
	"fmt"
	"sync"
	"sync/atomic"
)
//int64 is used with atomic

func main() {
	a := ("This is atomic/wait group...")
	var wg sync.WaitGroup
    var incrementer int64

	gs := 100
	wg.Add(gs)

	 fmt.Println("")
	 fmt.Println(a)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			r := atomic.LoadInt64(&incrementer)
			fmt.Println("")
			fmt.Println("-----")
			fmt.Println("# ", r)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("--------------------------------")
	fmt.Println("value:", incrementer)
	fmt.Println("--------------------------------")

}

