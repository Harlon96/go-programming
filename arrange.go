package main

import (
	"fmt"
	"runtime"
)

type person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := person{
		First: "Harlon",
		Last:  "Hutchison",
		Age:   43,
		Sayings: []string{
			"I love programming in my spare time",
			"Take it easy but be persistant.",
			"I want to go to sleep",
		},
	}
	u2 := person{
		First: "Brook",
		Last:  "Hutchison",
		Age:   43,
		Sayings: []string{
			"I love cook.",
			"I want to go to Ireland.",
			"I want to go to sleep also.",
		},
	}
	fmt.Println("")
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println(u1)
	fmt.Println("")
	fmt.Println("----------------------------------------------------------------------------")
	fmt.Println(u2)
	fmt.Println("")
	fmt.Println("") // this is space
	fmt.Println("")
	fmt.Println("CPU:", runtime.NumCPU)
	fmt.Println("")
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("")
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("")
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS)
}
