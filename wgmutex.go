package main

import(
	"fmt"
	"sync"
)

func main() {
    a := ("This is mutex/wait group...")
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println("")
			fmt.Println(a)
			fmt.Println("")
			fmt.Println("-----")
			fmt.Println("# ", incrementer)
			m.Unlock()
			wg.Done()
		}()
	}
		 wg.Wait()
		 fmt.Println("--------------------------------")
		 fmt.Println("value:", incrementer)
		 fmt.Println("--------------------------------")

}
