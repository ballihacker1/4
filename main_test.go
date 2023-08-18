package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
		{10, -5, 5},
		{100, 200, 300},
	}

	for _, c := range cases {
		result := Add(c.a, c.b)
		if result != c.expected {
			t.Errorf("Add(%d, %d) == %d, expected %d", c.a, c.b, result, c.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{10, 10, 0},
		{-5, 3, -8},
		{0, 5, -5},
	}

	for _, c := range cases {
		result := Subtract(c.a, c.b)
		if result != c.expected {
			t.Errorf("Subtract(%d, %d) == %d, expected %d", c.a, c.b, result, c.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{-2, 5, -10},
		{0, 100, 0},
		{7, 0, 0},
	}

	for _, c := range cases {
		result := Multiply(c.a, c.b)
		if result != c.expected {
			t.Errorf("Multiply(%d, %d) == %d, expected %d", c.a, c.b, result, c.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	cases := []struct {
		a, b         int
		expectedQuot int
		expectedErr  bool
	}{
		{10, 2, 5, false},
		{0, 5, 0, false},
		{8, 0, 0, true},
		{-15, 3, -5, false},
	}

	for _, c := range cases {
		result, err := Divide(c.a, c.b)

		if c.expectedErr {
			if err == nil {
				t.Errorf("Divide(%d, %d) expected an error, but got none", c.a, c.b)
			}
		} else {
			if err != nil {
				t.Errorf("Divide(%d, %d) encountered an unexpected error: %v", c.a, c.b, err)
			}

			if result != c.expectedQuot {
				t.Errorf("Divide(%d, %d) == %d, expected %d", c.a, c.b, result, c.expectedQuot)
			}
		}
	}
}
