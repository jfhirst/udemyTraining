package main

import (
	"fmt"
)

var z float64

func main() {

	x := 42
	y := 42.3445345343434543
	z = 35 / 6
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
	fmt.Printf("%T\n", z)
	fmt.Println(z)
}
