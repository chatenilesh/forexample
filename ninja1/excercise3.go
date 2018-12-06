package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true

	s := fmt.Sprintf("%d %s %t", x, y, z)
	fmt.Println(s)

	//x := fmt.Sprintf("%v %v %v" , x, y, z)
	x := fmt.Sprintf("%#v %#v %#v", x, y, z)
	fmt.Println(x)
}
