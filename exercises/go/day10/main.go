package main

import "fmt"

func main() {
	s := square{side: 4}
	c := circle{radius: 7}

	fmt.Println("Square")
	printArea(s)

	fmt.Println("Circle")
	printArea(c)
}
