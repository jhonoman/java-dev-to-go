package main

import "fmt"

func increment(x *int) {
	*x = *x + 1
}

func main() {
	n := 10
	increment(&n)
	fmt.Println(n)
}
