package main

import "fmt"

func main() {
	ret, err := fmt.Println("Hello World")
	fmt.Println(ret, err)

	ret1, err := fmt.Println("short declaration")
	fmt.Println(ret1, err)

	ret1, err = fmt.Println("no declaration")
	fmt.Println(ret1, err)

	ret1, _= fmt.Println("return throw away")
	fmt.Println(ret1)
}


//ebnf (Extended Backus-Naur form)
//identifier = letter {letter | unicode_digit}