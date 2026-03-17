package main

import "fmt"

// Quick analogy:
// goroutine = worker
// channel   = conveyor belt between workers
// make      = create the conveyor belt first

func producer(nums []int) <-chan int {
	out := make(chan int) // create a channel to send data
	go func() {          // goroutine that produces data into the channel asynchronously
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
	}()
	return out
}

func mapMultiply10(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * 10
		}
	}()
	return out
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	stream := producer(nums)
	evens := filterEven(stream)
	result := mapMultiply10(evens)

	fmt.Println("Channel pipeline result:")
	for n := range result {
		fmt.Println(n)
	}
}
