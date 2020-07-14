package main

import (
	"fmt"
	"time"
)

type MyError struct {
	Now time.Time
}

func (m *MyError) Error() string {
	fmt.Println("--- test before ---")

	// 这里会导致神奇的栈溢出
	// 由于此时m为nil, 会在里面触发 m 的error, 从而导致循环调用 error...  阔怕~
	// 由于此时的m是error类型, fmt打印的时候会调用 Error 方法, 从而实现了嵌套调用, 栈溢出
	str := fmt.Sprintf("error %v", m)

	// str += " test"
	fmt.Println("--- test after ---")

	return str
}

func Get() error {
	var m *MyError

	m = &MyError{
		time.Now(),
	}
	return m
}

func main() {

	var m *MyError
	str := fmt.Sprintf("hh %v", m)
	fmt.Println(str)

	return

	/*
	if e := Get(); e != nil {
		fmt.Println(e)
	}

	 */
}