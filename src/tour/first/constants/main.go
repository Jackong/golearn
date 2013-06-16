/**
 * Created with IntelliJ IDEA.
 * User: Jackong
 * Date: 13-6-12
 * Time: 下午3:28
 */
package main

import (
	"fmt"
)

const (
	Pi = 3.14
)

func needInt(x int) int {
	return x * 10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	const (
		World = "世界"
		Truth = true
		Big = 1 << 100
		Small = Big >> 99
	)
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules", Truth)

	fmt.Println(needFloat(Small))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
	//numeric constants are high-precision values
	//fmt.Println(needInt(Big))
}
