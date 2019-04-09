package main

import "fmt"

type A struct {
	val int
}

func newA(val int) (*A, *A) {
	r := &A{val: val}
	return r, r
}

func main() {
	prev, next := newA(100000)
	prev.val = 99
	fmt.Println(prev.val)
	fmt.Println(next.val)
}
