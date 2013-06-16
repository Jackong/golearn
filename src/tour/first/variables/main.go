/**
 * Created with IntelliJ IDEA.
 * User: Jackong
 * Date: 13-6-12
 * Time: 下午3:05
 */
package main

import (
	"fmt"
	"math/cmplx"
)

var x, y, z int
//Outside a function, := construct is not available
//p, q := 1, 3

func main() {
	//declares by var
	var c, python, java bool
	//with initializers
	var m, n, o = 1, false, "test"
	//short declarations
	a, b := 2, "nothing"
	var (
		Boolean = false
		MaxInt uint64 = 1<<64 - 1
		cmplex = cmplx.Sqrt(-5 + 21i)
	)

	fmt.Println(x, y, z, c, python, java)
	fmt.Println(m, n, o)
	fmt.Println(a, b)
	const (
		f = "%T(%v)\n"
	)
	fmt.Printf(f, Boolean, Boolean)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, cmplex, cmplex)
}
