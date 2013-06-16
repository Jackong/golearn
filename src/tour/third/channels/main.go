/**
 * User: Jackong
 * Date: 13-6-15
 * Time: 下午4:53
 */
package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int, t time.Duration) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	time.Sleep(t)
	c <- sum //send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[: len(a) / 2], c, 1000 * time.Millisecond)
	go sum(a[len(a) / 2 :], c, 1 * time.Millisecond)
	x, y := <-c, <-c //receive from c
	fmt.Println(x, y)

	c1 := make(chan int, 2)
	c1 <- 2
	c1 <- 3
	fmt.Println(<-c1)
	fmt.Println(<-c1)
	c1 <- 4
	c1 <- 3
	close(c1)
	v, ok := <-c1
	fmt.Println(v, ok)
	v, ok = <-c1
	fmt.Println(v, ok)
	//ok is false if there are no more values to receive and the channel is closed.
	v, ok = <-c1
	fmt.Println(v, ok)

	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	v, ok = <-c2
	fmt.Println(v, ok)
	for i := range c2 {
		fmt.Println(i)
	}
	v, ok = <-c2
	fmt.Println(v, ok) //ok is false
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) //close it for range to work normally
}
