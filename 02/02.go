package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := getInput("02.input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	twos, threes := ChecksumComponents(input)
	checksum := twos * threes
	fmt.Println(checksum)

	id1, id2, pos := FindNearlyIdenticalPair(input)
	letters := CommonLetters(id1, id2, pos)
	fmt.Println(letters)
}

func getInput(filename string) ([]string, error) {
	var strs []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return strs, nil
}

// ChecksumComponents takes an array of strings and returns the number of
// strings that contain two of any letter, and the number of strings that
// contain three of any letter.
func ChecksumComponents(strs []string) (int, int) {
	var twos, threes int

	for _, str := range strs {
		counts := make(map[rune]int)
		var hasTwo, hasThree bool

		for _, char := range str {
			counts[char]++
		}
		for _, count := range counts {
			if count == 2 {
				hasTwo = true
			} else if count == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}

	return twos, threes
}

// NearlyIdentical takes two strings of identical length and if their
// Hamming distance is one, returns the position of the differing letter,
// and -1 otherwise.
func NearlyIdentical(a string, b string) int {
	if len(a) != len(b) {
		return -1
	}
	pos := -1
	// don't range on the string as it returns runes, just check bytes
	for i, char := range []byte(a) {
		if char != b[i] {
			// not the first differing byte
			if pos > -1 {
				return -1
			}
			pos = i
		}
	}
	return pos
}

// FindNearlyIdenticalPair takes an array of strings, and returns a pair which
// have Hamming distance of one, or a pair of empty strings otherwise.
func FindNearlyIdenticalPair(input []string) (string, string, int) {
	for i, str := range input {
		for j := 0; j < len(input); j++ {
			// don't compare string against itself
			if i == j {
				continue
			}
			if pos := NearlyIdentical(str, input[j]); pos != -1 {
				return str, input[j], pos
			}
		}
	}
	return "", "", -1
}

// CommonLetters takes two strings with a Hamming distance of one, and returns
// a string that contains all letters in original sequence except the different
// letter.
func CommonLetters(a string, b string, pos int) string {
	var s strings.Builder
	s.WriteString(a[0:pos])
	s.WriteString(a[pos+1:])
	return s.String()
}
