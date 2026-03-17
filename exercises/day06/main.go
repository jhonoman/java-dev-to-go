package main

import "fmt"

func main() {
	scores := []int{78, 96, 68, 95, 77, 91, 65}
	total := 0
	for _, score := range scores {
		total += score
	}
	average := float64(total) / float64(len(scores))
	fmt.Printf("Average score: %.2f", average)
}
