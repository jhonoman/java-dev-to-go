# Daily Exercises + Concise Answers

Use this after trying on your own for at least 20-30 minutes.

## Day 1
**Question:** Create a program that prints:
- Hello, I am Budi
- I am 20 years old

**Answer:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, I am Budi")
    fmt.Println("I am 20 years old")
}
```

## Day 2
**Question:** Store your current age, then display your age 5 years from now.

**Answer:**
```go
package main

import "fmt"

func main() {
    age := 20
    fmt.Println("Current age:", age)
    fmt.Println("Age in 5 years:", age+5)
}
```

## Day 3
**Question:** Check the number `7`, then print "odd" or "even".

**Answer:**
```go
package main

import "fmt"

func main() {
    number := 7
    if number%2 == 0 {
        fmt.Println("even")
    } else {
        fmt.Println("odd")
    }
}
```

## Day 4
**Question:** Print numbers 1 to 10, then print only multiples of 3.

**Answer:**
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    fmt.Println("Multiples of 3:")
    for i := 1; i <= 10; i++ {
        if i%3 == 0 {
            fmt.Println(i)
        }
    }
}
```

## Day 5
**Question:** Create a function `divide(a, b int) (int, error)` that is safe when `b == 0`.

**Answer:**
```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divisor must not be zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

## Day 6
**Question:** Create a slice of scores and calculate the average.

**Answer:**
```go
package main

import "fmt"

func main() {
    scores := []int{80, 90, 75, 85}
    total := 0
    for _, s := range scores {
        total += s
    }
    average := float64(total) / float64(len(scores))
    fmt.Println("Average:", average)
}
```

## Day 7
**Question:** Combine simple (hardcoded) input data, process it, and print a report.

**Concise answer:**
- Use a simple struct (name + score)
- Loop through a student slice
- Calculate total and average

## Day 8
**Question:** Create a `Student` struct with an `Introduce()` method.

**Answer:**
```go
package main

import "fmt"

type Student struct {
    Name string
    Age  int
}

func (s Student) Introduce() {
    fmt.Printf("Hello, I am %s, and I am %d years old\n", s.Name, s.Age)
}

func main() {
    s := Student{Name: "Budi", Age: 20}
    s.Introduce()
}
```

## Day 9
**Question:** Change a variable value using a pointer in a function.

**Answer:**
```go
package main

import "fmt"

func incrementByOne(x *int) {
    *x = *x + 1
}

func main() {
    n := 10
    incrementByOne(&n)
    fmt.Println(n)
}
```

## Day 10
**Question:** Create a `Shape` interface with an `Area()` method.

**Concise answer:**
- Define the `Shape` interface
- Implement it for `Square` and `Circle`
- Create a function `printArea(s Shape)`

## Day 11
**Question:** Create an age validation function (minimum age: 17).

**Concise answer:**
- If `<17`, return `error`
- If valid, return `nil`
- Handle it in `main` with `if err != nil`

## Day 12
**Question:** Split simple math functions into a `utils` package.

**Concise answer:**
- Run `go mod init ...`
- Put exported function(s) (for example `Add`) in `utils/math.go`
- Import the package in `main.go`

## Day 13
**Question:** Add tests for function `Add(a, b int)`.

**Concise answer:**
- Create `math_test.go`
- Use table-driven tests
- Run `go test ./...`

## Day 14
**Question:** Create a mini CLI: add todo items, list all items.

**Concise answer:**
- Store data in-memory (slice)
- Use `os.Args` for command handling
- Minimum commands: `add`, `list`

---
If you want, the next step is I can create **14 separate exercise files** (`day-01` to `day-14`), each ready to run with `go run` and with the question as a top comment.
