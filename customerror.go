package main

import(
	"fmt"
)

type customErr struct {
	info string
	aged float64
}

func (ce customErr) Error() string {
	  return fmt.Sprintf(`Here is the error: %v %v`, ce.info, ce.aged)
}

func main() {
      c1 := customErr{
      	   info: `I am in dire need of some caffeine`,
      	   aged: 3.14,
	  }
	  do(c1)
}

func do(e error) {
	  fmt.Println(``)
	  fmt.Println(`  Do it, you can do it!`, e, ``, e.(customErr).info,``, e.(customErr).aged)
}