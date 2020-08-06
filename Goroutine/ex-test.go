package main

import (
	"fmt"
)

const (
	INT_MAX = int(^uint(0) >> 1)
	// INT_MIN = ^INT_MAX
	INT_MIN = 1 << 63
)

func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func main() {
	/*t := ""
	arr := strings.Split(t, "&")
	fmt.Println(len(arr))
	fmt.Println(arr[0])
	query := "a"
	subt  := ""
	pos := strings.Index(query, subt)
	fmt.Println(pos)

	fmt.Println(len(""))
	log.Fatal()
	*/

	/*
	buf := make([]byte, 100)
	itoa(&buf, INT_MAX, 1)
	*/

	// fmt.Println(string(buf))
	fmt.Println(INT_MIN)

	// sync.WaitGroup{}
}
