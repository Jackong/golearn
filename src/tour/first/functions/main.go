/**
 * Created with IntelliJ IDEA.
 * User: Jackong
 * Date: 13-6-12
 * Time: 下午2:13
 */
package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

/*
also parameters can share the same type
 */
func add2(x, y int) int {
	return x + y
}

func add3(x string, y, z int) int {
	return y + z
}

func add4(x, y int, z string) int {
	return x + y
}

//-----------multiple results-----------------

func swap(x, y string) (string, string) {
	return y, x
}

//------------named results-------------------
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//-------------function values----------------
var hypot = func(x, y int) int {
	return x + y
}

var fv = split

//-------------function closures---------------
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println(add(12, 22))
	fmt.Println(add2(12, 22))
	fmt.Println(add3("nothing", 12, 22))
	fmt.Println(add4(12, 22, "nothing"))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(19))

	fmt.Println(fv(19))
	fmt.Println(hypot(11, 22))

	pos, neg := adder(), adder()
	fmt.Println(pos(1), neg(-2))
	fmt.Println(pos(2), neg(-4))
	fmt.Println(pos(5), neg(-7))
}


