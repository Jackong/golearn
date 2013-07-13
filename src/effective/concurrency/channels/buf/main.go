/**
 * User: Jackong
 * Date: 13-6-29
 * Time: 上午9:02
 */
package main

import (
	"fmt"
	"time"
)


func main() {
	complete := make(chan int, 5)

	go func() {
		fmt.Println("send a signal and unblocking")
		time.Sleep(10 * 60)
		complete <- 1
	}()

	fmt.Println("receiver always blocking to get a signal")
	signal := <-complete
	fmt.Println("get the signal", signal)

	fmt.Println()

	fmt.Println("bufferd channel can send a signal and unblocking")
	complete <- 2

	signal = <-complete
	fmt.Println("get the signal", signal)
}




