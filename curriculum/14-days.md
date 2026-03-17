# 14-Day Go Learning Curriculum

## Week 1 — Language Foundations

### Day 1 — Setup & Hello Go
- Topic: Go installation, `package main`, `func main()`, `fmt.Println`
- Practice: print your name, age, and city
- Outcome: can compile and run your first program

### Day 2 — Variables & Data Types
- Topic: `var`, `:=`, `string`, `int`, `bool`, simple conversions
- Practice: calculate your age 5 years from now
- Outcome: understand declarations and basic types

### Day 3 — Operators & Conditionals
- Topic: arithmetic operators, comparisons, `if/else`
- Practice: check whether a number is even or odd
- Outcome: can implement decision-making logic

### Day 4 — Loops
- Topic: classic `for`, `for range`, `break`, `continue`
- Practice: print numbers 1-20 and multiples of 3
- Outcome: can iterate through data

### Day 5 — Basic Functions
- Topic: parameters, return values, multiple return values
- Practice: rectangle area function and safe division function
- Outcome: can break logic into functions

### Day 6 — Array, Slice, Map
- Topic: array vs slice differences, append operations, map key/value
- Practice: store student scores and calculate the average
- Outcome: understand core Go data structures

### Day 7 — Weekly Review
- Topic: recap Day 1-6
- Practice: combined mini exercise (input + process + output)
- Outcome: strong foundation before OOP-style patterns in Go

## Week 2 — Go Program Structure

### Day 8 — Struct & Method
- Topic: `type`, `struct`, receiver methods
- Practice: `Student` struct with `Introduce()` method
- Outcome: can model data and behavior

### Day 9 — Basic Pointers
- Topic: `&` and `*`, pointers in functions
- Practice: change a variable value through a function
- Outcome: understand pass-by-value and references

### Day 10 — Basic Interface
- Topic: interface as a behavior contract
- Practice: `Shape` with `Area()` for square and circle
- Outcome: understand Go-style polymorphism

### Day 11 — Error Handling
- Topic: `if err != nil` pattern, `errors.New`, `%w` wrapping
- Practice: validate age input
- Outcome: can write clear and meaningful errors

### Day 12 — Package & Module
- Topic: `go mod init`, local packages, export vs unexport
- Practice: split calculation utilities into `utils` package
- Outcome: can organize projects more cleanly

### Day 13 — Basic Testing
- Topic: `testing` package, simple table-driven tests
- Practice: test calculator functions
- Outcome: understand Go testing habits

### Day 14 — Concurrency + Mini Project
- Topic: goroutines, channels (intro)
- Practice: simple todo/balance mini CLI
- Outcome: complete a final mini project and identify next steps

## Next Steps (After 14 Days)
- Build one real project: simple API using `net/http`
- Learn database access (`database/sql`) and JSON
- Learn lightweight clean architecture
