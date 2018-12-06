package main

import (
	"fmt"
)

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	const (
		_   = iota
		ikb = 1 << (iota * 10)
		imb = 1 << (iota * 10)
		igb = 1 << (iota * 10)
	)

	fmt.Printf("%d\t\t\t%b\n", ikb, ikb)
	fmt.Printf("%d\t\t\t%b\n", imb, imb)
	fmt.Printf("%d\t\t%b\n", igb, igb)
}
