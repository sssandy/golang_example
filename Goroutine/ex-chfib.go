package main

import (
	"fmt"
	"os"
)

func fibonacci(n int, ch chan int) {
	x := 0
	y := 1
	for i := 0; i < n; i++ {
		ch <- x
		// 太简洁了 牛批
		x, y = y, x + y
	}

	close(ch)
}

func main() {
	ch := make(chan int, 10)

	go fibonacci(cap(ch), ch)

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(len("str"))
	// fmt.Println(cap(s))

	f, _ := os.Open("aa")
	f.Close()
}
