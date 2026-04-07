package main

import (
	"fmt"
	"day12/utils"
)

func main() {
	a := 10
	b := 5
	fmt.Println("Addition:", utils.Add(a, b))
	fmt.Println("Subtraction:", utils.Subtract(a, b))
	fmt.Println("Multiplication:", utils.Multiply(a, b))
	result, err := utils.Divide(a, b)
	if err != nil {
		fmt.Println("Division:", err)
	} else {
		fmt.Println("Division:", result)
	}
}