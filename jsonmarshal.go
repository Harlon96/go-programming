package main

import(
	"encoding/json"
	"errors"
    "fmt"
)
type person struct {
	First string
	Last string
	Sayings []string
	err error
}

func main() {
     p1 := person{
     	  First: "Harlon",
     	  Last: "Hutch",
     	  Sayings: []string{`I am doing well`, `I am doing great thus far`},
	 }

	 p2 := person{
	 	First: "gabe",
	 	Last: "forfeather",
	 	Sayings: []string{`Life is like a box of chocolates`},
	 }

	 bs, err := json.Marshal(p1)
	 if err != nil {
	 	fmt.Println(err) //you can use log.Println(err)
		 //errors.New(fmt.Sprintf(`there was an error is json: %v`, err)) will work here but not best practice
	 	return
	 }
	 fmt.Println(string (bs))

	cs, err := json.Marshal(p2)
	if err != nil {
		fmt.Println(err) //you can use log.Println(err)
		//errors.New(fmt.Sprintf(`there was an error is json: %v`, err)) will work here but not best practice
		return
	}
	fmt.Println(string (cs))
}

func git(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf(`there was an error is json: %v`, err))
	}
	return bs, nil
}

func got(b interface{}) ([]byte, error) {
	cs, err := json.Marshal(b)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf(`there was an error is json: %v`, err))
	}
	return cs, nil
}