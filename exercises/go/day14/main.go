// Jalankan:
//   go run . demo             — demo dasar goroutine & channel
//   go run . add Belajar HTTP — tambah todo baru
//   go run . list             — tampilkan semua todo (via goroutine + channel)

package main

import (
	"fmt"
	"os"
	"strings"
)

// ── Data in-memory ────────────────────────────────────────────────────────────

var todos = []string{
	"Learn goroutine",
	"Learn channel",
	"Build a mini project",
}

// ── Goroutine + Channel ───────────────────────────────────────────────────────

// worker receives todo items from the jobs channel and prints them one by one.
// Once the jobs channel is closed, the goroutine sends a done signal.
func worker(jobs <-chan string, done chan<- bool) {
	i := 1
	for item := range jobs { // ranging over a channel stops automatically when it is closed
		fmt.Printf("  %d. %s\n", i, item)
		i++
	}
	done <- true // notify main that the goroutine has finished
}

// demoConcurrency shows the most basic goroutine + channel usage.
func demoConcurrency() {
	// make(chan T) — creates an unbuffered channel of type string
	ch := make(chan string)

	// Two goroutines run concurrently…
	go func() {
		ch <- "Hello from the first goroutine"  // send message
	}()
	go func() {
		ch <- "Hello from the second goroutine" // send message
	}()

	// …main receives both (order may differ each run)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("Both goroutines finished!")
}

// ── CLI ───────────────────────────────────────────────────────────────────────

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := strings.ToLower(os.Args[1])

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . add <todo text>")
			return
		}
		item := strings.Join(os.Args[2:], " ")
		todos = append(todos, item)
		fmt.Println("Todo added:", item)
		fmt.Println("(Data is in-memory only — resets on next run)")

	case "list":
		if len(todos) == 0 {
			fmt.Println("No todos yet.")
			return
		}

		// Send each todo to the worker goroutine via a channel
		jobs := make(chan string)
		done := make(chan bool)

		go worker(jobs, done) // run worker in a separate goroutine

		fmt.Println("Todo List (printed by worker goroutine):")
		for _, item := range todos {
			jobs <- item // send item to channel
		}
		close(jobs) // close channel → worker goroutine knows there are no more items

		<-done // wait for goroutine to finish before exiting

	case "demo":
		fmt.Println("=== Goroutine & Channel Demo ===")
		demoConcurrency()

	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("=== Day 14: Concurrency + Mini CLI Todo ===")
	fmt.Println("Commands:")
	fmt.Println("  go run . add <todo text>  — add a new todo")
	fmt.Println("  go run . list             — list all todos (goroutine + channel)")
	fmt.Println("  go run . demo             — basic goroutine & channel demo")
}
