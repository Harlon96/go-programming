package main

import(
	"fmt"
	"sort"
)

type person struct {
	First string
	Last string
	age int
}

type ByFirst []person

func (bn ByFirst) Len() int           { return len(bn) }
func (bn ByFirst) Swap(i, j int)      { bn[i],bn[j] = bn[j], bn[i] }
func (bn ByFirst) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := person{"harlon","hutch",43}
	p2 := person{"brook","hutch",43}
	p3 := person{"trevor","chip",12}
	p4 := person{"kendra","jax",13}
	p5 := person{"tiff","bake",41}

	people := []person{p1, p2, p3, p4, p5}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("----------------------------")
	fmt.Println("out of alphabetical order...")
	fmt.Println("----------------------------")
	fmt.Println("")
	fmt.Println(people)
	fmt.Println("")
	fmt.Println("------------------------")
	fmt.Println("in alphabetical order...")
	fmt.Println("------------------------")
	fmt.Println("")
	fmt.Println("")
	sort.Sort(ByFirst(people))
	fmt.Println(people)
}
