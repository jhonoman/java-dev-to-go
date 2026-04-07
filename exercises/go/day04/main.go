package main

import "fmt"

func main() {
	number := 20
	for i := 1; i <= number; i++ {
		if i != number {
			fmt.Print(i, ",")
		} else {
			fmt.Print(i)
		}
	}
	fmt.Println()
	for i := 1; i <= number; i++ {
		if i%3 == 0 && i+3 < number {
			fmt.Print(i, ",")
		} else if i%3 == 0 && i+3 >= number {
			fmt.Print(i)
		}
	}
}
