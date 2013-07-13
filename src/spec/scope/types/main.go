/**
 * User: Jackong
 * Date: 13-6-19
 * Time: 下午11:18
 */
package main

import "fmt"

type Mutex struct {
	Aa int
	Bb string
}
func (m *Mutex) Lock() {}
func (m Mutex) Unlock() {}
type NewMutex Mutex

type PtrMutex *Mutex

type PrintableMutex struct {
	Mutex
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

//interface: has the same method set as Block
type MyBlock Block
func main() {
	//no any method
	//nm := NewMutex{1, "2"}
	//pm := PtrMutex{}

	//has all method and anonymous field Mutex
	//pm2 := PrintableMutex{}

	var x interface{} = 7  // x has dynamic type int and value 7
	i, ok := x.(int)           // i has type int and value 7
	fmt.Println(i, ok)
}
