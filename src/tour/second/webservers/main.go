/**
 * User: Jackong
 * Date: 13-6-15
 * Time: 下午4:27
 */
package main

import (
	"fmt"
	"net/http"
)

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello~")
}

func main() {
	var (
		h Hello
	)
	http.ListenAndServe("localhost:4000", h)
}
