package main

import "fmt"

func main() {
	//exercise 1
	a := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a)

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", a)

	//exercise 2
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(s)

	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", s)

	//exercise 3
	s2 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(s2)
	fmt.Println(s2[:5])
	fmt.Println(s2[5:])
	fmt.Println(s2[2:7])
	fmt.Println(s2[1:6])

	//exercise 4
	s3 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s3 = append(s3, 52)
	fmt.Println(s3)

	s3 = append(s3, 53, 54, 55)
	fmt.Println(s3)

	x := []int{56, 57, 58, 59, 60}
	s3 = append(s3, x...)
	fmt.Println(s3)

	//exercise 5
	s4 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s4 = append(s4[:3], s4[6:]...)
	fmt.Println(s4)

	//exercise 6
	states := make([]string, 5, 5)
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	//exercise 7
	ms1 := []string{"James", "Bond", "Shaken, not stirred"}
	ms2 := []string{"Miss", "Moneypenny", "Hellooooo, James"}

	//ms3 := [][]string{{"A"}, {"B"}}
	ms3 := [][]string{ms1, ms2}
	fmt.Println(ms3)

	for _, xs := range ms3 {
		//for i, xs := range ms3 {
		//fmt.Println("Outer ", i, xs)
		fmt.Println("Outer ", xs)
		for j, ys := range xs {
			fmt.Println("Inner ", j, ys)
		}
	}

	//exercise 8
	ppl := map[string][]string{
		"bond_james": []string{"Shaken, not stirred", "Martinis", "Women"},
		"no_dr":      []string{"Be evil", "Ice cream", "Sunsets"},
	}
	fmt.Println(ppl)

	//exercise 9
	ppl["added_key"] = []string{"test1", "test2"}
	for k, v := range ppl {
		//fmt.Println(k, v)
		fmt.Println("Map: ", k)
		for i, s := range v {
			fmt.Println(i, s)
		}
	}

	//exercise 10
	delete(ppl, "added_key")
	for k, v := range ppl {
		//fmt.Println(k, v)
		fmt.Println("Map: ", k)
		for i, s := range v {
			fmt.Println(i, s)
		}
	}
}
