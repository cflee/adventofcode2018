package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	lines, err := getInput("04.input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sleepiness := ProcessGuardSleepiness(lines)
	sleepiestGuard, sleepiestMin := ComputeMostSleepyGuard(sleepiness)
	fmt.Println(sleepiestGuard * sleepiestMin)

	sleepiestMomentGuard, sleepiestMomentMin := ComputeMostSleepyMoment(sleepiness)
	fmt.Println(sleepiestMomentGuard * sleepiestMomentMin)
}

// This returns the input, sorted in lexicographic ascending order (which is
// increasing time, due to the YYYY-MM-DD HH:MM format)
func getInput(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// lexicographic order for 'yyyy-mm-dd hh:mm' is ascending order of time
	sort.Strings(lines)

	return lines, nil
}

// ProcessGuardSleepiness takes a list of ordered-by-time event strings and
// computes for each guard, the number of shifts in which they were asleep
// during each minute during the midnight hour (:00-:59).
// The input list must be in ascending timestamp order.
// Since there is always exactly one guard on duty each night (starting from
// midnight), we can assume the shift change always happens at the first minute
// (:00) of the hour, and can ignore the specific timestamp beyond using it
// for ordering.
// Since the sleep/wake actions are always within the midnight hour, first
// sleep is always within the timeframe and we don't need to care about a guard
// starting a shift asleep.
func ProcessGuardSleepiness(lines []string) map[int][60]int {
	sleepiness := make(map[int][60]int)
	guard := 0
	start := 0

	for _, l := range lines {
		if strings.HasSuffix(l, "begins shift") {
			// trim off the timestamp in front
			fmt.Sscanf(l[19:], "Guard #%d begins shift", &guard)
		} else if strings.HasSuffix(l, "falls asleep") {
			fmt.Sscanf(l[15:17], "%d", &start)
		} else if strings.HasSuffix(l, "wakes up") {
			var end int
			fmt.Sscanf(l[15:17], "%d", &end)
			for i := start; i < end; i++ {
				g := sleepiness[guard]
				g[i] = g[i] + 1
				sleepiness[guard] = g
			}
		}
	}

	return sleepiness
}

// ComputeMostSleepyGuard takes a map of guard IDs to how many shifts they have
// been sleeping during each minute, and returns the ID of the guard that spends
// the most time sleeping, and the minute in which they are most commonly
// sleeping
func ComputeMostSleepyGuard(sleepiness map[int][60]int) (int, int) {
	sleepiestGuard := 0
	sleepiestSum := 0
	sleepiestMin := 0

	for guard, counts := range sleepiness {
		// I suppose we could do some fancy sorting, but hey, one loop through..
		guardSum := 0
		guardMaxCount := 0
		guardMaxMin := 0

		for min, c := range counts {
			guardSum += c
			if c > guardMaxCount {
				guardMaxCount = c
				guardMaxMin = min
			}
		}

		if guardSum > sleepiestSum {
			sleepiestGuard = guard
			sleepiestSum = guardSum
			sleepiestMin = guardMaxMin
		}
	}
	return sleepiestGuard, sleepiestMin
}

// ComputeMostSleepyMoment takes a map of guard IDs to how many shifts they have
// been sleeping during each minute, and returns the ID of the guard that is
// most frequently sleeping on the same minute, and the specific minute.
func ComputeMostSleepyMoment(sleepiness map[int][60]int) (int, int) {
	sleepiestGuard := 0
	sleepiestCount := 0
	sleepiestMin := 0

	for guard, counts := range sleepiness {
		// I suppose we could do some fancy sorting, but hey, one loop through..
		guardMaxCount := 0
		guardMaxMin := 0

		for min, c := range counts {
			if c > guardMaxCount {
				guardMaxCount = c
				guardMaxMin = min
			}
		}

		if guardMaxCount > sleepiestCount {
			sleepiestGuard = guard
			sleepiestCount = guardMaxCount
			sleepiestMin = guardMaxMin
		}
	}
	return sleepiestGuard, sleepiestMin
}
