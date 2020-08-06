package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right,ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		v2, ok := <-ch2
		if ok == false {
			return false
		}

		if i != v2 {
			return false
		}
	}

	return true
}

func main() {
	t := tree.New(1)
	ch := make(chan int)

	go Walk(t, ch)

	for cv := range ch {
		fmt.Println(cv)
	}

	// fmt.Println(<-ch)
	cv, ok := <-ch
	fmt.Println(cv, ok)


	// s := Same(tree.New(1), tree.New(1))
	// fmt.Println(s)

}
