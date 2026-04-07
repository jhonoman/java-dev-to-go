package test

import (
	"testing"

	"day12/utils"
)

type intCase struct {
	name string
	a    int
	b    int
	want int
}

func runIntCases(t *testing.T, cases []intCase, fn func(int, int) int, label string) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := fn(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("%s(%d, %d) = %d, want %d", label, tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	cases := []intCase{
		{name: "positive", a: 10, b: 5, want: 15},
		{name: "negative", a: -3, b: -7, want: -10},
		{name: "mixed", a: -3, b: 7, want: 4},
	}
	runIntCases(t, cases, utils.Add, "Add")
}

func TestSubtract(t *testing.T) {
	cases := []intCase{
		{name: "positive", a: 10, b: 5, want: 5},
		{name: "negative result", a: 5, b: 10, want: -5},
		{name: "with negative", a: -3, b: 7, want: -10},
	}
	runIntCases(t, cases, utils.Subtract, "Subtract")
}

func TestMultiply(t *testing.T) {
	cases := []intCase{
		{name: "positive", a: 4, b: 5, want: 20},
		{name: "multiply by zero", a: 4, b: 0, want: 0},
		{name: "negative", a: -4, b: 5, want: -20},
	}
	runIntCases(t, cases, utils.Multiply, "Multiply")
}

func TestDivide(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		got, err := utils.Divide(10, 2)
		if err != nil {
			t.Fatalf("Divide(10, 2) unexpected error: %v", err)
		}
		if got != 5 {
			t.Fatalf("Divide(10, 2) = %d, want %d", got, 5)
		}
	})

	t.Run("divide by zero", func(t *testing.T) {
		_, err := utils.Divide(10, 0)
		if err == nil {
			t.Fatal("Divide(10, 0) expected error, got nil")
		}
	})

	t.Run("pembulatan int", func(t *testing.T) {
		got, err := utils.Divide(7, 2)
		if err != nil {
			t.Fatalf("Divide(7, 2) unexpected error: %v", err)
		}
		if got != 3 {
			t.Fatalf("Divide(7, 2) = %d, want %d", got, 3)
		}
	})
}
