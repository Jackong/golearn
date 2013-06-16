/**
 * Created with IntelliJ IDEA.
 * User: Jackong
 * Date: 13-6-13
 * Time: 下午8:25
 */
package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 3})

	v := Vertex{2, 4}
	v.X = 4
	fmt.Println(v.X)

	q := &v
	q.Y = 1e9
	fmt.Println(v)

	var (
		a = Vertex{1, 2}
		b = &Vertex{1, 2}
		c = Vertex{Y:1}
		d = Vertex{}
	)
	fmt.Println(a, b, c, d)

	var k *Vertex = new(Vertex)
	//var k = new(Vertex)
	//k := new(Vertex)
	fmt.Println(k)
	k.X, k.Y = 11, 2
	fmt.Println(k)
}
