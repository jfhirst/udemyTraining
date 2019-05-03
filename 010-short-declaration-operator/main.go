package main

import "fmt"

var z = 123
var b int
var c string
var d bool

func main() {
	//   short declaration operator
	// DECLARE a variable and assign a value to it
	// works fiune inside a function e.g  func main {}
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)

	// works if define outside a fuction
	var y = 43
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	foo()
}

func foo() {

	fmt.Println("again", z)
}
