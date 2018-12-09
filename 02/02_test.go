package main

import "testing"

func TestChecksumComponents(t *testing.T) {
	testCases := []struct {
		desc      string
		input     []string
		expected1 int
		expected2 int
	}{
		{
			desc:      "Example",
			input:     []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"},
			expected1: 4,
			expected2: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual1, actual2 := ChecksumComponents(tC.input); actual1 != tC.expected1 || actual2 != tC.expected2 {
				t.Errorf("Expected checksum values for %v to be %d, %d but instead got %d, %d", tC.input, tC.expected1, tC.expected2, actual1, actual2)
			}
		})
	}
}

func TestNearlyIdentical(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []string
		expected int
	}{
		{
			desc:     "",
			input:    []string{"abcde", "axcye"},
			expected: -1,
		},
		{
			desc:     "",
			input:    []string{"fghij", "fguij"},
			expected: 2,
		},
		{
			desc:     "",
			input:    []string{"abcd", "abc"},
			expected: -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := NearlyIdentical(tC.input[0], tC.input[1]); actual != tC.expected {
				t.Errorf("Expected %v to be %d but instead got %d", tC.input, tC.expected, actual)
			}
		})
	}
}
