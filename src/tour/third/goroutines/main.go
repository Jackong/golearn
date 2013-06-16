/**
 * User: Jackong
 * Date: 13-6-15
 * Time: 下午4:43
 */
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("World")
	say("Hello")
}
