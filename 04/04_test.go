package main

import (
	"reflect"
	"testing"
)

func TestProcessGuardSleepiness(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []string
		expected map[int][60]int
	}{
		{
			desc: "",
			input: []string{
				"[1518-11-01 00:00] Guard #10 begins shift",
				"[1518-11-01 00:05] falls asleep",
				"[1518-11-01 00:25] wakes up",
				"[1518-11-01 00:30] falls asleep",
				"[1518-11-01 00:55] wakes up",
				"[1518-11-01 23:58] Guard #99 begins shift",
				"[1518-11-02 00:40] falls asleep",
				"[1518-11-02 00:50] wakes up",
				"[1518-11-03 00:05] Guard #10 begins shift",
				"[1518-11-03 00:24] falls asleep",
				"[1518-11-03 00:29] wakes up",
				"[1518-11-04 00:02] Guard #99 begins shift",
				"[1518-11-04 00:36] falls asleep",
				"[1518-11-04 00:46] wakes up",
				"[1518-11-05 00:03] Guard #99 begins shift",
				"[1518-11-05 00:45] falls asleep",
				"[1518-11-05 00:55] wakes up",
			},
			expected: map[int][60]int{
				10: [60]int{
					0, 0, 0, 0, 0, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 2, 1, 1, 1, 1, 0,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
				99: [60]int{
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
					2, 2, 2, 2, 2, 3, 2, 2, 2, 2,
					1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := ProcessGuardSleepiness(tC.input); !reflect.DeepEqual(actual, tC.expected) {
				t.Errorf("Expected sleepiness for %v to be %v but instead got %v", tC.input, tC.expected, actual)
			}
		})
	}
}

func TestComputeMostSleepyGuard(t *testing.T) {
	testCases := []struct {
		desc          string
		input         map[int][60]int
		expectedGuard int
		expectedMin   int
	}{
		{
			desc: "",
			input: map[int][60]int{
				10: [60]int{
					0, 0, 0, 0, 0, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 2, 1, 1, 1, 1, 0,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
				99: [60]int{
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
					2, 2, 2, 2, 2, 3, 2, 2, 2, 2,
					1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			},
			expectedGuard: 10,
			expectedMin:   24,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actualGuard, actualMin := ComputeMostSleepyGuard(tC.input); actualGuard != tC.expectedGuard || actualMin != tC.expectedMin {
				t.Errorf("Expected guard/min for %v to be %d/%d but instead got %d/%d", tC.input, tC.expectedGuard, tC.expectedMin, actualGuard, actualMin)
			}
		})
	}
}

func TestComputeMostSleepyMoment(t *testing.T) {
	testCases := []struct {
		desc          string
		input         map[int][60]int
		expectedGuard int
		expectedMin   int
	}{
		{
			desc: "",
			input: map[int][60]int{
				10: [60]int{
					0, 0, 0, 0, 0, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 2, 1, 1, 1, 1, 0,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
					1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
				99: [60]int{
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
					0, 0, 0, 0, 0, 0, 1, 1, 1, 1,
					2, 2, 2, 2, 2, 3, 2, 2, 2, 2,
					1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			},
			expectedGuard: 99,
			expectedMin:   45,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actualGuard, actualMin := ComputeMostSleepyMoment(tC.input); actualGuard != tC.expectedGuard || actualMin != tC.expectedMin {
				t.Errorf("Expected guard/min for %v to be %d/%d but instead got %d/%d", tC.input, tC.expectedGuard, tC.expectedMin, actualGuard, actualMin)
			}
		})
	}
}
