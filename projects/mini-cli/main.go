package main

import (
    "encoding/json"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

const dataFile = "todos.json"

type Todo struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Done  bool   `json:"done"`
}

var todos []Todo

func main() {
    if err := loadTodos(); err != nil {
        fmt.Println("Failed to load todos:", err)
        return
    }

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
        addTodo(strings.Join(os.Args[2:], " "))
    case "list":
        listTodos()
    case "done":
        if len(os.Args) < 3 {
            fmt.Println("Usage: go run . done <todo id>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid todo id:", os.Args[2])
            return
        }
        markDone(id)
    case "delete":
        if len(os.Args) < 3 {
            fmt.Println("Usage: go run . delete <todo id>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid todo id:", os.Args[2])
            return
        }
        deleteTodo(id)
    case "clear":
        clearTodos()
    case "help":
        printHelp()
    default:
        printHelp()
    }
}

func loadTodos() error {
    f, err := os.Open(dataFile)
    if err != nil {
        if os.IsNotExist(err) {
            todos = []Todo{}
            return nil
        }
        return err
    }
    defer f.Close()

    bytes, err := io.ReadAll(f)
    if err != nil {
        return err
    }

    if len(bytes) == 0 {
        todos = []Todo{}
        return nil
    }

    if err := json.Unmarshal(bytes, &todos); err != nil {
        return err
    }
    return nil
}

func saveTodos() error {
    bytes, err := json.MarshalIndent(todos, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(dataFile, bytes, 0644)
}

func addTodo(title string) {
    title = strings.TrimSpace(title)
    if title == "" {
        fmt.Println("Todo text cannot be empty.")
        return
    }

    todo := Todo{
        ID:    nextID(),
        Title: title,
        Done:  false,
    }
    todos = append(todos, todo)

    if err := saveTodos(); err != nil {
        fmt.Println("Failed to save todos:", err)
        return
    }
    fmt.Printf("Todo added: [%d] %s\n", todo.ID, todo.Title)
}

func listTodos() {
    if len(todos) == 0 {
        fmt.Println("No todos yet.")
        return
    }

    fmt.Println("Todo List:")
    for _, todo := range todos {
        status := "[ ]"
        if todo.Done {
            status = "[x]"
        }
        fmt.Printf("%s %d. %s\n", status, todo.ID, todo.Title)
    }
}

func markDone(id int) {
    idx := findTodoIndexByID(id)
    if idx == -1 {
        fmt.Println("Todo not found.")
        return
    }

    todos[idx].Done = true
    if err := saveTodos(); err != nil {
        fmt.Println("Failed to save todos:", err)
        return
    }
    fmt.Printf("Todo marked done: [%d] %s\n", todos[idx].ID, todos[idx].Title)
}

func deleteTodo(id int) {
    idx := findTodoIndexByID(id)
    if idx == -1 {
        fmt.Println("Todo not found.")
        return
    }

    removed := todos[idx]
    todos = append(todos[:idx], todos[idx+1:]...)
    if err := saveTodos(); err != nil {
        fmt.Println("Failed to save todos:", err)
        return
    }
    fmt.Printf("Todo deleted: [%d] %s\n", removed.ID, removed.Title)
}

func clearTodos() {
    todos = []Todo{}
    if err := saveTodos(); err != nil {
        fmt.Println("Failed to save todos:", err)
        return
    }
    fmt.Println("All todos cleared.")
}

func nextID() int {
    maxID := 0
    for _, t := range todos {
        if t.ID > maxID {
            maxID = t.ID
        }
    }
    return maxID + 1
}

func findTodoIndexByID(id int) int {
    for i, t := range todos {
        if t.ID == id {
            return i
        }
    }
    return -1
}

func printHelp() {
    fmt.Println("Mini CLI Todo")
    fmt.Println("Commands:")
    fmt.Println("  go run . add <todo text>")
    fmt.Println("  go run . list")
    fmt.Println("  go run . done <todo id>")
    fmt.Println("  go run . delete <todo id>")
    fmt.Println("  go run . clear")
    fmt.Println("  go run . help")
}
