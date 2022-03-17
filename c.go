package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)
	if x%3 == 0 {
		fmt.Println("fizz")
	} else if x%5 == 0 {
		fmt.Println("Bazz")
	}
}
