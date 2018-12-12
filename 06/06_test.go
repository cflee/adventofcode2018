package main

import "testing"

func TestFindBoundaries(t *testing.T) {
	testCases := []struct {
		desc      string
		input     []Coordinate
		expected1 Coordinate
		expected2 Coordinate
	}{
		{
			desc: "",
			input: []Coordinate{
				Coordinate{x: 1, y: 1},
				Coordinate{x: 1, y: 6},
				Coordinate{x: 8, y: 3},
				Coordinate{x: 3, y: 4},
				Coordinate{x: 5, y: 5},
				Coordinate{x: 8, y: 9},
			},
			expected1: Coordinate{x: 1, y: 1},
			expected2: Coordinate{x: 8, y: 9},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual1, actual2 := FindBoundaries(tC.input); actual1 != tC.expected1 || actual2 != tC.expected2 {
				t.Errorf("Expected bounds for %v to be %v/%v but instead got %v/%v", tC.input, tC.expected1, tC.expected2, actual1, actual2)
			}
		})
	}
}

func TestDist(t *testing.T) {
	testCases := []struct {
		desc     string
		input1   Coordinate
		input2   Coordinate
		expected int
	}{
		{
			desc:     "",
			input1:   Coordinate{x: 8, y: 9},
			input2:   Coordinate{x: 1, y: 6},
			expected: 10,
		},
		{
			desc:     "",
			input1:   Coordinate{x: 1, y: 6},
			input2:   Coordinate{x: 8, y: 9},
			expected: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := Dist(tC.input1, tC.input2); actual != tC.expected {
				t.Errorf("Expected dist for %v/%v to be %d but instead got %d", tC.input1, tC.input2, tC.expected, actual)
			}
		})
	}
}

func TestBruteforceLargestArea(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []Coordinate
		expected int
	}{
		{
			desc: "",
			input: []Coordinate{
				Coordinate{x: 1, y: 1},
				Coordinate{x: 1, y: 6},
				Coordinate{x: 8, y: 3},
				Coordinate{x: 3, y: 4},
				Coordinate{x: 5, y: 5},
				Coordinate{x: 8, y: 9},
			},
			expected: 17,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := BruteforceLargestArea(tC.input); actual != tC.expected {
				t.Errorf("Expected largest area for %v to be %d but instead got %d", tC.input, tC.expected, actual)
			}
		})
	}
}
