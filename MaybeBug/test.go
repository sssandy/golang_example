package main

import (
	"fmt"
)

func Check(b []byte) {
	fmt.Printf("0x%p \n", b)
	b = append(b, 'A', 'B')
	fmt.Printf("0x%p \n", b)
	// fmt.Println("b:" + string(b[0]))
	// b[0] = 'A'
}

func main() {
	/*
	var b []byte
	fmt.Println(cap(b))
	fmt.Printf("0x%p \n", &b)

	b = append(b, 'A')
	fmt.Printf("0x%p \n", &b)

	fmt.Println(b[0])
	*/

	var c []byte
	// var b []byte
	c = make([]byte, 0, 1)
	// c = c[:1]
	// c = append(c , 'A')
	fmt.Printf("0x%p \n", &c)

	Check(c)

	c = c[:1]
	fmt.Println(c, len(c), cap(c))
}
