package main

import "fmt"

type student struct {
	name  string
	age   int
	class string
}

func introduce(s student) {
	fmt.Printf("Name: %s, Age: %d, Class: %s\n", s.name, s.age, s.class)
}

func main() {
	students := []student{
		{"Budi", 20, "12A"},
		{"Siti", 19, "11B"},
		{"Andi", 21, "12C"},
	}

	for _, s := range students {
		introduce(s)
	}
}
