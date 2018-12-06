package main

import "fmt"

type nyc int

func main() {
	var x nyc

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
