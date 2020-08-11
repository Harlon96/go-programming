package main

import(
	"encoding/json"
	"fmt"
)
//Unmarshal

type person struct {
	 First string   `json:"First"`
	 Last string    `json:"Last"`
	 Age int        `json:"Age"`
}

func main() {
	     a := "Having some fun with json/Unmarshal..."
	     s := `[{"First":"Harlo","Last":"Michael","Age":43},{"First":"Brook","Last":"Candy","Age":27}]`
	     bs := []byte(s)

	     fmt.Println("")
	     fmt.Println(a)
	     fmt.Println("")
	     fmt.Println("--------------------------------------")
	     fmt.Printf("%T\n",a)
	     fmt.Printf("%T\n",s)
	     fmt.Printf("%T\n", bs)
         fmt.Println("--------------------------------------")
	     fmt.Println("")

	     var people []person

	     err := json.Unmarshal(bs, &people)
	      if err != nil {
	      	  fmt.Println(err)
		  }
		  fmt.Println("\nHere is the two people", people)
	      fmt.Println("")

	     for i, v := range people {
	     	   fmt.Println("\nPerson Number", i)
			   fmt.Println("")
	     	   fmt.Println(v.First, v.Last, v.Age)
		 }
}
