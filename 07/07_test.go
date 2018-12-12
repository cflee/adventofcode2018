package main

import "testing"

func TestSortSteps(t *testing.T) {
	testCases := []struct {
		desc     string
		input1   map[string][]string
		input2   map[string][]string
		expected string
	}{
		{
			desc: "",
			input1: map[string][]string{
				"A": []string{"C"},
				"F": []string{"C"},
				"B": []string{"A"},
				"D": []string{"A"},
				"E": []string{"B", "D", "F"},
			},
			input2: map[string][]string{
				"C": []string{"A", "F"},
				"A": []string{"B", "D"},
				"B": []string{"E"},
				"D": []string{"E"},
				"F": []string{"E"},
			},
			expected: "CABDFE",
		},
		{
			desc: "Sample swapped D and F",
			input1: map[string][]string{
				"A": []string{"C"},
				"F": []string{"A"},
				"B": []string{"A"},
				"D": []string{"C"},
				"E": []string{"B", "D", "F"},
			},
			input2: map[string][]string{
				"C": []string{"A", "D"},
				"A": []string{"B", "F"},
				"B": []string{"E"},
				"D": []string{"E"},
				"F": []string{"E"},
			},
			expected: "CABDFE",
		},
		{
			desc: "Two starting points C and G",
			input1: map[string][]string{
				"B": []string{"C"},
				"Y": []string{"C"},
				"D": []string{"G"},
				"Z": []string{"D", "Y"},
			},
			input2: map[string][]string{
				"C": []string{"B", "Y"},
				"Y": []string{"Z"},
				"G": []string{"D"},
				"D": []string{"Z"},
			},
			expected: "CBGDYZ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := SortSteps(tC.input1, tC.input2); actual != tC.expected {
				t.Errorf("Expected order for %v/%v to be %s but instead got %s", tC.input1, tC.input2, tC.expected, actual)
			}
		})
	}
}

func TestSortStepsWithTime(t *testing.T) {
	testCases := []struct {
		desc      string
		in        map[string][]string
		out       map[string][]string
		workerNum int
		stepDelay int
		expected  int
	}{
		{
			desc: "",
			in: map[string][]string{
				"A": []string{"C"},
				"F": []string{"C"},
				"B": []string{"A"},
				"D": []string{"A"},
				"E": []string{"B", "D", "F"},
			},
			out: map[string][]string{
				"C": []string{"A", "F"},
				"A": []string{"B", "D"},
				"B": []string{"E"},
				"D": []string{"E"},
				"F": []string{"E"},
			},
			workerNum: 2,
			stepDelay: 0,
			expected:  15,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := SortStepsWithTime(tC.in, tC.out, tC.workerNum, tC.stepDelay); actual != tC.expected {
				t.Errorf("Expected order for %v/%v/%d/%d to be %d but instead got %d", tC.in, tC.out, tC.workerNum, tC.stepDelay, tC.expected, actual)
			}
		})
	}
}
