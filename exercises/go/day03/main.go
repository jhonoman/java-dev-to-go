package main

import "fmt"

func main() {
	number := 9
	if number%2 == 0 {
		fmt.Println(number, "is an even number")
	} else {
		fmt.Println(number, "is an odd number")
	}
}
