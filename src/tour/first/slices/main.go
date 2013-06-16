/**
 * Created with IntelliJ IDEA.
 * User: Jackong
 * Date: 13-6-13
 * Time: 下午8:40
 */
package main

import (
	"fmt"
)


func main() {
	p := []int{2, 3, 4, 6, 7, 13, 44}
	fmt.Println(p)
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] = %d\n", i, p[i])
	}

	//slicing slices
	fmt.Println(p[1:4])
	fmt.Println(p[:3])
	fmt.Println(p[4:])
	fmt.Println(p[4])
	fmt.Println(p[4:4])
	fmt.Printf("%d---%d\n",len(p), cap(p))
	//TODO: why?
	fmt.Println(p[7:7])
	//fmt.Println(p[8:8])
	//fmt.Println(p[7:6])

	//making slices
	a := make([]int, 5)
	printSlice(a)
	b := make([]int, 0, 5)
	printSlice(b)
	c := b[:2]
	printSlice(c)
	d := c[2:5]
	printSlice(d)

	//nil slices
	var z []int
	fmt.Println(z, len(z), cap(z))
	if nil == z {
		fmt.Println("nil")
	}

	//range
	pow := []int{1,2,3,4,5,6,7,8}
	for i, v := range pow {
		fmt.Printf("pow[%d] = %d\n", i, v)
	}
	for i := range pow{
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	//len and cap
	k := []int{0,1,2,3,4,5}
	o := k[1:5]
	printSlice(o)//len = 4 cap = 5
	x := k[2:4]
	printSlice(x)//len = 2 cap = 4
	x = x[1:cap(x)]
	printSlice(x)//len = 3 cap = 3
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}
