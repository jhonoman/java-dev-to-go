package main

import "fmt"

type studentScore struct {
	name  string
	score int
}

func main() {
	students := []studentScore{
		{"Budi", 78},
		{"Andi", 96},
		{"Siti", 68},
		{"Rina", 95},
		{"Dedi", 77},
		{"Tina", 91},
		{"Joko", 65},
	}
	total := students[0].score
	highest := students[0]
	lowest := students[0]
	for _, s := range students {
		total += s.score
		if s.score > highest.score {
			highest = s
		} else if s.score < lowest.score {
			lowest = s
		}
	}
	average := float64(total) / float64(len(students))
	fmt.Printf("Average score: %.2f\n", average)
	fmt.Printf("Highest score: %s with score %d\n", highest.name, highest.score)
	fmt.Printf("Lowest score: %s with score %d\n", lowest.name, lowest.score)
}
