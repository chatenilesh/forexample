package main

import "fmt"

var y = 43

var z int // default value assigned

func main() {
	ret, err := fmt.Println("Hello World")
	fmt.Println(ret, err)

	ret1, err := fmt.Println("short declaration")
	fmt.Println(ret1, err)

	ret1, err = fmt.Println("no declaration")
	fmt.Println(ret1, err)

	ret1, _ = fmt.Println("return throw away")
	fmt.Println(ret1)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	foo()

	fmt.Println(z)
}

func foo() {
	z = 1
}

//ebnf (Extended Backus-Naur form)
//identifier = letter {letter | unicode_digit}
