package main;

import "fmt";

func main() {
	fmt.Println("Hi");
	foo();
	fmt.Println("After foo");

	for i:=0;i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ");
		}
	}

	bar();
}

func foo() {
	fmt.Println("I am in foo");
}


func bar() {
	fmt.Println("\nAnd then we exited")
}