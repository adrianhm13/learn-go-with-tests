package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("got %d, expected %d given, %v", got, expected, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !slices.Equal(got, expected) {
			t.Errorf("got %d, expected %d", got, expected)
		}
	}

	t.Run("make the sum of multiple slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 5}, []int{0, 1, 2})
		expected := []int{5, 3}

		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 5, 5})
		expected := []int{0, 10}

		checkSums(t, got, expected)
	})
}
