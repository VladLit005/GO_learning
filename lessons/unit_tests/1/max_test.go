package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	testTable := []struct {
		numbers  []int
		expected int
	}{
		{
			numbers:  []int{1, 2, 3, -4, 5, 70, -3, 4},
			expected: 70,
		},
		{
			numbers:  []int{-4, -3, -74},
			expected: 0,
		},
		{
			numbers:  []int{},
			expected: 0,
		},
	}
	for _, v := range testTable {

		result := Max(v.numbers)

		t.Logf("Calling Max(%v), result %d", v.numbers, result)

		if result != v.expected {
			t.Errorf("Incorrect result. Expect %d, got %d", v.expected, result)
		}
	}

}

func TestMaxIndex(t *testing.T) {
	numbers := []int{1, 2, 3, -4, 5, 70, -3, 4}
	expected := 5

	result := MaxIndex(numbers)

	if result != expected {
		t.Errorf("Incorrect result. Expect %d, got %d", expected, result)
	}
}
