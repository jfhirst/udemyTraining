package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t  %x \t %#X %q \n", i, i, i, i, i)
		//tab is \t
		//  # shows it as hex  eg 0x2A
	}
}
