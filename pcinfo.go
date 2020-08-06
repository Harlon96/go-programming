package main

import(
      "fmt"
      "runtime"
)

// lets you see your OS, ARCH, and CPU type

func main() {

	x := string(`Your pc's information is as follows:`)

	fmt.Println(x)
	fmt.Println("-----------------------------------")
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU's\t\t", runtime.NumCPU())
	fmt.Println("-----------------------------------")
}
