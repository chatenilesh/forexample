package main

import "fmt"

func main() {

	//limit scope of x
	if x := 42; x == 42 {
		fmt.Println("001")
	}
	//fmt.Println(x)

}
