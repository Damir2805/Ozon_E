package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	n := 10
	
	sum(c3, c1, c2, n, func(x int) int { return x + x })
	go enumerate(c1, 1)
	go enumerate(c2, 2)
	
	for i := 0; i < n; i++ {
		fmt.Println(<-c3)
	}
}

func sum(out chan<- int, c1 <-chan int, c2 <-chan int, n int, f func(int) int) {
	go func() {
		for i := 0; i < n; i++ {
			x := <- c1
			y := <- c2
			out <- f(x) + f(y)
		}
	}()
}

func enumerate(out chan<- int, start int)  {
	for i := start; ; i++ {
		out <- i
	}
}
