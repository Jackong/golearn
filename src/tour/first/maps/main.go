/**
 * Created with IntelliJ IDEA.
 * User: Jackong
 * Date: 13-6-13
 * Time: 下午9:27
 */
package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var (
	m map[string]Vertex
)
func main() {
	m = make(map[string]Vertex)
	m["Bell"] = Vertex{
		40.111,
		-74.156654,
	}
	fmt.Println(m["Bell"])
	fmt.Println(m)

	var m2 = map[string]Vertex {
		"Bele": Vertex {
			123.46,
			4561.5464,
		},
		"Google": Vertex {
			4564.4543,
			213.21463,
		},
	}
	fmt.Println(m2)

	m3 := map[string]Vertex {
		"Bex": {124, 16},
		"Gox": {},
	}
	fmt.Println(m3)

	delete(m3, "Bex")
	fmt.Println(m3)
	fmt.Println(m3["Bex"])
	fmt.Println(m3)

	delete(m3, "Goxxx")
	fmt.Println(m3)

	v, ok := m3["Gox"]
	fmt.Println(v, ok)
	v, ok = m3["Bex"]
	fmt.Println(v, ok)
}

