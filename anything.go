package main

import "fmt"

func main() {
	ret, err := fmt.Println("Hello World")
	fmt.Println(ret, err)

	ret1, _ := fmt.Println("Hello World")
	fmt.Println(ret1)
}