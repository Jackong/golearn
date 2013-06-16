/**
 * User: Jackong
 * Date: 13-6-15
 * Time: 下午5:32
 */
package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		fmt.Println("/3.x")
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	fmt.Println("/1")
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("/4.x")
			fmt.Println(<-c)
		}
		quit <- 0
		fmt.Println("/5")
	}()
	fmt.Println("/2")
	fibonacci(c, quit)

	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
