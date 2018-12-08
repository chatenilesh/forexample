package main

import "fmt"

func main() {
	//exercise 1
	for i := 1; i <= 10000; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	//exercise 2
	for c := int('A'); c <= int('Z'); c++ {
		fmt.Println(c - int('A') + 1)
		for i := 0; i < 3; i++ {
			//fmt.Printf("\t%#U '%c'\n", byte(c), c)
			fmt.Printf("\t%#U\n", byte(c))
		}
		fmt.Println()
	}

	//exercise 3
	yr := 0
	for yr <= 1990 {
		fmt.Println(yr)
		yr++
	}

	//exercise 4
	yr = 1990
	for {
		if yr > 2018 {
			break
		}

		fmt.Println(yr)
		yr++
	}

	//exercise 5
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}

	//exercise 6
	if true {
		fmt.Println("in if")
	}

	//exercise 7
	x := 50
	if x < 10 {
		fmt.Println("in if")
	} else if x > 10 && x < 40 {
		fmt.Println("in else if ")
	} else {
		fmt.Println("in else")
	}

	//exercise 8
	switch {
	case 1 < 1:
		fmt.Println("1<1")
	case 1 == 3:
		fmt.Println("1==3")
	default:
		fmt.Println("default")
	}

	//exercise 9
	s := "favSport"
	switch s {
	case "running":
		fmt.Println("running")
	case "favSport":
		fmt.Println("favSport")
	default:
		fmt.Println("default")
	}

	//exercise 9
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
