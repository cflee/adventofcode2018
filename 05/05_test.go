package main

import "testing"

func TestReact(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected int
	}{
		{
			desc:     "",
			input:    "dabAcCaCBAcCcaDA",
			expected: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := React(tC.input); len(actual) != tC.expected {
				t.Errorf("Expected reacted length for %s to be %d but instead got %d", tC.input, tC.expected, len(actual))
			}
		})
	}
}

func TestShortestReactWithRemoval(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected int
	}{
		{
			desc:     "",
			input:    "dabAcCaCBAcCcaDA",
			expected: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := ShortestReactWithRemoval(tC.input); len(actual) != tC.expected {
				t.Errorf("Expected reacted length for %s to be %d but instead got %d", tC.input, tC.expected, len(actual))
			}
		})
	}
}
