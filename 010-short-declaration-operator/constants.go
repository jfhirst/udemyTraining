package main

import (
	"fmt"
)

const a = 42
const b = 42.78
const c = "James Bond"

const (
	d int     = 99
	e float64 = 99.99
	f string  = "J. Bond"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)

}
