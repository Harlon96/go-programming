package main

import(
     "fmt"
     "runtime"
)

func main() {
       go foo()
       bar()
}

func foo() {
         fmt.Println("")
         fmt.Println("----------------")
         fmt.Println("OS:\t",runtime.GOOS)
         fmt.Println("CPU:\t",runtime.NumCPU())
         fmt.Println("ARCH:\t",runtime.GOARCH)
         fmt.Println("Goroutines:", runtime.NumGoroutine())
         fmt.Println("----------------")
         fmt.Println("")

         for i := 0; i < 15; i++ {

             fmt.Println("foo count:", i)
       }
}


func bar() {
         fmt.Println("")
         fmt.Println("----------------")
         fmt.Println("OS:\t",runtime.GOOS)
         fmt.Println("CPU:\t",runtime.NumCPU())
         fmt.Println("ARCH:\t",runtime.GOARCH)
         fmt.Println("Goroutine:",runtime.NumGoroutine())
         fmt.Println("----------------")
         fmt.Println("")

          for j := 0; j < 15; j++ {

             fmt.Println("bar count:", j)
      }
}
