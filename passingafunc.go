package main

import (
	"fmt"
)

//some func options
//passing a func

//even though the code passes through.. it does not work properly.

func main() {
        li := []int{0,1,2,3,5,8,13,21,34, 55}
        s := sum(li...)
        a := (`passing a func`)


	    fmt.Println("")
	    fmt.Println("")
	    fmt.Println("---------------")
	    fmt.Println("")
        fmt.Println(a)
	    fmt.Println("")
        fmt.Println("---------------")
        fmt.Println("")
        fmt.Println("correct output:", s)
	    fmt.Println("")
	    fmt.Println("---------------")



      //  s2 := even(sum, li...)
     //   fmt.Println("even numbers:")

      //  s3 := odd(sum, li...)
      //  fmt.Println("odd numbers:")
}
func sum(li ...int) int{
	 total := 0
	 for _, v := range li {
	 	 total += v
	 }

	 for s := 0; s < 15; s++ {
		 fmt.Println("")
		 fmt.Println("")
		 fmt.Println("---------------")
		 fmt.Println("")
		 fmt.Println("output should be 142:", s)
	}
     return total


}