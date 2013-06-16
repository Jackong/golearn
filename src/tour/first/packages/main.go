/**
 * Created with IntelliJ IDEA.
 * User: Jackong
 * Date: 13-6-12
 * Time: 下午1:52
 */
package main

import (
	"fmt"
	"math"
)
/*
//also you can import them like following
import "fmt"
import math "math" //rename the package
*/

func main()	{
	fmt.Println("Happy", math.Pi, "Day")

	//a name is exported if it begins with a capital letter.
	//fmt.Println("Happy", math.pi, "Day")
}
