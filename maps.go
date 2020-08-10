package main

import(
	"fmt"
)

//this is a map data structure program

func main() {
	s := ("I am counting eleven of the states...")
    q := ("--------------------------------------")
    x := make([]string, 11, 11)
    x = []string{`Alabama`,`Alaska`,`Arizona`,`Arkansas`,`California`,`Colorado`,`Connecticut`,`Delaware`,`Florida`,`Georgia`,`Hawaii`}
    jb := []string{`James Bond`, `Shaken not stirred`} //slice of a string
    mp := []string{`Harlo`, `Money pants`} //slice of a string
    xp := [][]string{jb,mp} //string of a string

	m := map[string]int {  //map data structure
	    "Harlo": 43,
		"Sassy": 37,
		"Morpheus": 45,
}

fmt.Println(``)              // m := map printing it out
	fmt.Println(q)
fmt.Println(m)
	fmt.Println(``)
	fmt.Println(q)
fmt.Println(m[`Harlo`])
fmt.Println(m[`Sassy`])
fmt.Println(m[`Morpheus`])

	fmt.Println(``)

m[`hutch`] = 128                 // printing out the value of m[hutch] if v is equal to ok

	fmt.Println(``)
	fmt.Println(q)

if v, ok := m[`hutch: 128`]; ok {    //making hutch = 128 and printing it out
fmt.Println(v)
}

	fmt.Println(``)
	fmt.Println(q)

for k, v := range m {     //changing m to equal k, v loop
fmt.Println(k, v)
}
fmt.Printf("\n")
fmt.Println("")

tiff := []int{0,1,2,3,5,8,13,21,34}         // for range tiff loop
for i, v := range tiff {
fmt.Println(i,v)
	fmt.Println(``)
	fmt.Println(q)
fmt.Println(`count: fibonacci numbers`)
	fmt.Println(``)
	fmt.Println(q)
}
//printing out eleven of the united states alphabetically

 fmt.Println(s)
	fmt.Println(q)
 fmt.Println(x)
	fmt.Println(q)
	fmt.Println(``)
 fmt.Println(len(x))
 fmt.Println(cap(x))

	fmt.Println("")
	fmt.Println("")


for i := 0; i < len(x); i++ {
	 fmt.Println(i, x[i])
}

//multidimensional slices being printed out with a for range loop

    fmt.Println("")
	fmt.Println("")

for i, xs := range xp {
	fmt.Println("record #:", i) // first range loop

	for j, val := range xs {
		fmt.Printf("\t index position: %v \t value: \t %v \n", j, val) //second range loop
	}
}

  fmt.Println(jb)
  fmt.Println(q)
  fmt.Println(mp)
  fmt.Println(q)
  fmt.Println(xp)

}


