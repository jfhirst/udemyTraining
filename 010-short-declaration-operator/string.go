package main

import (
	"fmt"
)

func main() {
	s := "my string"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	const sample2 = "\u2318"

	fmt.Println(sample)
	fmt.Println(sample2)

	bs := []byte(s)
	fmt.Println(bs)
	//fmt.Printf("%T\n" ,bs)
	//fmt.Printf("%x\n" ,bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
		//	fmt.Printf("%#x \n", s[i])

	}

	fmt.Println("\n")
	for i, v := range s {
		fmt.Println(i, v)
		fmt.Printf("at index position %d we have hex %#X", i, v)
	}
}
