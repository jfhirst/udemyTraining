package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	z := runtime.NumCPU
	fmt.Printf("%d \t %b \t  %x \t %#X \n", z, z, z, z)

}
