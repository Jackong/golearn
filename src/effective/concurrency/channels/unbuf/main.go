/**
 * User: Jackong
 * Date: 13-6-29
 * Time: 上午8:43
 */
package main

import (
	"fmt"
)

func main() {
	complete := make(chan int)

	go func() {
		fmt.Println("send a signal and blocking for 'complete'")
		complete <- 1
		fmt.Println("unblock complete-sender")
	}()

	fmt.Println("blocking to get a signal")
	signal := <-complete
	fmt.Println("get the signal", signal)

	fmt.Println()

/*	fmt.Println("unbuffered channel can send a signal and blocking")
	complete <- 2
	fmt.Println("blocking, so it is a wrong code")

	signal = <-complete
	fmt.Println("get the signal", signal)*/
}
