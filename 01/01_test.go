package main

import "testing"

func TestFinalFreq(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{1, -2, 3, 1}, 3},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
	}

	for _, test := range tests {
		if actual := FinalFreq(test.input); actual != test.expected {
			t.Errorf("Expected the final freq of %v to be %d but instead got %d!", test.input, test.expected, actual)
		}
	}
}

func TestFirstRepeatedFreq(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{1, -2, 3, 1}, 2},
		// {[]int{1, -1}, 0},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, 3, 8, 5, -6}, 5},
		{[]int{7, 7, -2, -7, -4}, 14},
	}

	for _, test := range tests {
		if actual := FirstRepeatedFreq(test.input); actual != test.expected {
			t.Errorf("Expected the final freq of %v to be %d but instead got %d!", test.input, test.expected, actual)
		}
	}

}
