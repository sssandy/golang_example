package main

func fib(ch, quit chan int) {
	x := 0
	y := 1

	select {
		case <-ch:
			x, y = y, x + y
		case <-quit:
			return
	}

	return
}

func main() {
	ch := make(chan int)
	quit

}
