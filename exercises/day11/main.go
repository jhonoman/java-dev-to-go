package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) error {
	if age < 17 {
		return errors.New("you are not old enough")
	}
	return nil
}

func main() {
	age := 10
	if err := checkAge(age); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("You are old enough")
}
