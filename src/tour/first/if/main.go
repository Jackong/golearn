/**
 * Created with IntelliJ IDEA.
 * User: Jackong
 * Date: 13-6-12
 * Time: 下午3:46
 */
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		return lim - v
	}
	//can't use v here
	return lim// - v
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
