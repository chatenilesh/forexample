package main

import "fmt"

func main() {
	a := 7
	b := 42

	fmt.Println(a <= b)
	fmt.Println(a == b)
	fmt.Println(a >= b)

	//uint8  ~ byte
	//uint32 ~ rune

	s := "Hello World"
	//s [5] = '-' NOT allowed
	//s:="Hello, 世界"
	fmt.Println(s, len(s))
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println()

	for i, v := range s {
		fmt.Printf("%d %T %d\n", i, v, v)
	}
}
