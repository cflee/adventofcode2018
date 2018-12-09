package main

import (
	"testing"
)

func TestBuildClaimGrid(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []Claim
		expected map[Coordinate][]int
	}{
		{
			desc: "",
			input: []Claim{
				{id: 123, c: Coordinate{x: 3, y: 2}, w: 4, h: 2},
			},
			expected: map[Coordinate][]int{
				Coordinate{x: 3, y: 2}: []int{123},
				Coordinate{x: 4, y: 2}: []int{123},
				Coordinate{x: 5, y: 2}: []int{123},
				Coordinate{x: 6, y: 2}: []int{123},
				Coordinate{x: 3, y: 3}: []int{123},
				Coordinate{x: 4, y: 3}: []int{123},
				Coordinate{x: 5, y: 3}: []int{123},
				Coordinate{x: 6, y: 3}: []int{123},
			},
		},
		{
			desc: "",
			input: []Claim{
				{id: 123, c: Coordinate{x: 1, y: 2}, w: 1, h: 1},
			},
			expected: map[Coordinate][]int{
				Coordinate{x: 1, y: 2}: []int{123},
			},
		},
		{
			desc: "",
			input: []Claim{
				{id: 124, c: Coordinate{x: 2, y: 1}, w: 1, h: 1},
				{id: 123, c: Coordinate{x: 2, y: 1}, w: 1, h: 1},
			},
			expected: map[Coordinate][]int{
				Coordinate{x: 2, y: 1}: []int{124, 123},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := BuildClaimGrid(tC.input)
			if len(actual) != len(tC.expected) {
				t.Errorf("Expected grid for %v to have %d items but instead got %d items", tC.input, len(tC.expected), len(actual))
			}
			for k, v := range tC.expected {
				if len(v) != len(actual[k]) {
					t.Errorf("Expected grid for %v to have %d at %v but instead got %d", tC.input, v, k, actual[k])
				}
			}
			// if reflect.DeepEqual(actual, tC.expected) {
			// 	t.Errorf("Expected grid for %v to be %v but instead got %v", tC.input, tC.expected, actual)
			// }
		})
	}
}

func TestCountClaimGridOverlaps(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []Claim
		expected int
	}{
		{
			desc: "",
			input: []Claim{
				{id: 1, c: Coordinate{x: 1, y: 3}, w: 4, h: 4},
				{id: 2, c: Coordinate{x: 3, y: 1}, w: 4, h: 4},
				{id: 3, c: Coordinate{x: 5, y: 5}, w: 2, h: 2},
			},
			expected: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := CountClaimGridOverlaps(BuildClaimGrid(tC.input)); actual != tC.expected {
				t.Errorf("Expected overlaps for %v to be %d but instead got %d", tC.input, tC.expected, actual)
			}
		})
	}
}

func TestFindViableClaim(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []Claim
		expected int
	}{
		{
			desc: "",
			input: []Claim{
				{id: 1, c: Coordinate{x: 1, y: 3}, w: 4, h: 4},
				{id: 2, c: Coordinate{x: 3, y: 1}, w: 4, h: 4},
				{id: 3, c: Coordinate{x: 5, y: 5}, w: 2, h: 2},
			},
			expected: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := FindViableClaim(BuildClaimGrid(tC.input)); actual != tC.expected {
				t.Errorf("Expected claim id for %v to be %d but instead got %d", tC.input, tC.expected, actual)
			}
		})
	}
}
