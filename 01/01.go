package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := getInput("01.input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(FinalFreq(input))
	fmt.Println(FirstRepeatedFreq(input))
}

func getInput(filename string) ([]int, error) {
	var nums []int

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, i)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}

// FinalFreq takes an array of integer frequency changes, and returns the final
// frequency from a starting frequency of 0.
func FinalFreq(changes []int) int {
	freq := 0
	for _, val := range changes {
		freq += val
	}
	return freq
}

// FirstRepeatedFreq takes an array of integer frequency changes, and returns
// the first frequency that is seen a second time.
func FirstRepeatedFreq(changes []int) int {
	freq := 0
	seen := make(map[int]bool)
	for {
		for _, val := range changes {
			freq += val
			if seen[freq] {
				return freq
			}
			seen[freq] = true
		}
	}
}
