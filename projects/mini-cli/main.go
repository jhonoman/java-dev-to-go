package main

import (
    "fmt"
    "os"
    "strings"
)

var todos []string

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
    case "list":
        if len(todos) == 0 {
            fmt.Println("No todos yet.")
            return
        }
        fmt.Println("Todo List:")
        for i, item := range todos {
            fmt.Printf("%d. %s\n", i+1, item)
        }
    default:
        printHelp()
    }
}

func printHelp() {
    fmt.Println("Mini CLI Todo")
    fmt.Println("Commands:")
    fmt.Println("  go run . add <todo text>")
    fmt.Println("  go run . list")
}
