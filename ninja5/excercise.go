package main

import "fmt"

type person struct {
	firstName          string
	lastName           string
	favIceCreamFlavour []string
}

func main() {
	//exercise 1

	p1 := person{
		firstName:          "a",
		lastName:           "b",
		favIceCreamFlavour: []string{"q", "w"},
	}

	p2 := person{
		firstName:          "c",
		lastName:           "d",
		favIceCreamFlavour: []string{"x", "y"},
	}

	fmt.Println(p1.firstName, p1.lastName)
	fmt.Println(p1)
	fmt.Println(p2)

	for k, v := range p1.favIceCreamFlavour {
		fmt.Println(k, v)
	}

	//exercise 2
	//ps := map[string]person{
	//				p1.lastName: p1,
	// 			}
	ps := map[string]person{}
	ps[p1.lastName] = p1
	ps[p2.lastName] = p2

	fmt.Println(ps)

	for k, v := range ps {
		fmt.Println(k, v.firstName)
	}

	for k, v := range ps["d"].favIceCreamFlavour {
		fmt.Println(k, v)
	}

	//exercise 3
	type vechile struct {
		doors string
		color string
	}

	type truck struct {
		vechile
		fourWheel bool
	}

	type sedan struct {
		vechile
		luxury bool
	}

	trk1 := truck{
		vechile: vechile{
			doors: "doorv",
			color: "color",
		},
		fourWheel: false,
	}

	fmt.Println(trk1)
	fmt.Println(trk1.color)
	fmt.Println(trk1.doors)
	fmt.Println(trk1.fourWheel)

	sdn1 := sedan{
		vechile: vechile{
			doors: "doorv",
			color: "color",
		},
		luxury: true,
	}

	fmt.Println(sdn1)
	fmt.Println(sdn1.color)
	fmt.Println(sdn1.doors)
	fmt.Println(sdn1.luxury)

	//exercise 4
	anyn1 := struct {
		name  string
		cars  map[string]int
		house []string
	}{
		name: "anynonmous struct",
		cars: map[string]int{
			"c1": 1,
			"c2": 2,
			"c3": 3,
		},
		house: []string{"h1", "h2"},
	}

	fmt.Println(anyn1)
	fmt.Println(anyn1.name)
	fmt.Println(anyn1.cars["c3"])
	fmt.Println(anyn1.cars["c4"])
	fmt.Println(anyn1.house[1])

	for k, v := range anyn1.cars {
		fmt.Println(k, v)
	}

	for k, v := range anyn1.house {
		fmt.Println(k, v)
	}
}
