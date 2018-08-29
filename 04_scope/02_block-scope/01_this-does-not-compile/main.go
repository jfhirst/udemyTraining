package main

import "fmt"
   //var x = 75

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// no access to x
	// this does not compile
	//x := 77
	fmt.Println(x)
}
