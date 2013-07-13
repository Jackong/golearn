/**
 * User: Jackong
 * Date: 13-6-19
 * Time: 下午10:55
 */
package main

import (
	"fmt"
)

const (
	One01 = 2
	One02
	One03
)

const (
	Zero1 = iota
	One1
	Two1
)
const (
	Zero2 = 2
	One2 = iota
	Two2
	Three2
)

const (
	Zero3 = 5
	One3 = 3
	Two3 = iota
	Three3
)

const (
	_ = iota
	One
	Two
	Three
)

const (
	EST  = (1 + iota)
	CST
	MST
	PST
)

func main() {
	fmt.Println(One01)
	fmt.Println(One02)
	fmt.Println(One03)

	fmt.Println(Zero1)
	fmt.Println(One1)
	fmt.Println(Two1)

	fmt.Println(Zero2)
	fmt.Println(One2)
	fmt.Println(Two2)
	fmt.Println(Three2)

	fmt.Println(Zero3)
	fmt.Println(One3)
	fmt.Println(Two3)
	fmt.Println(Three3)

	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)

	fmt.Println(EST)
	fmt.Println(CST)
	fmt.Println(MST)
	fmt.Println(PST)
}
