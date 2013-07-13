/**
 * User: Jackong
 * Date: 13-7-10
 * Time: 下午8:15
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("test1", 0777)
	os.MkdirAll("test1/test2/test3", 0777)
	err := os.Remove("test1")
	fmt.Println(err)
	os.RemoveAll("test1")
}
