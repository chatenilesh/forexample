package main

import "fmt"

const (
	a     = 42
	b int = 43
)

func main() {
	fmt.Println(a, b)

	c := a << 1
	fmt.Printf("%d %b %x\n", c, c, c)

	roster_l := `asd
				asd
				asdasd
				sad`
	fmt.Println(roster_l)

	const (
		w = 2019 + iota
		x = 2019 + iota
		y = 2019 + iota
		z = 2019 + iota
	)

	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
