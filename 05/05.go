package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	lines, err := getInput("05.input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reacted := React(lines)
	fmt.Println(len(reacted))

	shortestWithRemoval := ShortestReactWithRemoval(reacted)
	fmt.Println(len(shortestWithRemoval))
}

func getInput(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return line, nil
}

func invertCase(r rune) rune {
	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}
	return unicode.ToLower(r)
}

// React takes a polymer and returns the length of the fully reacted polymer.
func React(polymer string) string {
	s := NewStack()
	for _, u := range polymer {
		if s.Size() > 0 && s.Peek() == invertCase(u) {
			s.Pop()
		} else {
			s.Push(u)
		}
	}
	return s.String()
}

// ShortestReactWithRemoval takes a polymer and returns the shortest polymer
// that could arise by reacting the polymer after removing all of one unit type.
func ShortestReactWithRemoval(polymer string) string {
	attempted := make(map[string]bool)
	shortest := ""

	for _, r := range polymer {
		lower := string(unicode.ToLower(r))
		upper := string(unicode.ToUpper(r))
		if attempted[lower] == false {
			attempted[lower] = true
			stripped := strings.Replace(strings.Replace(polymer, lower, "", -1), upper, "", -1)
			result := React(stripped)
			if shortest == "" || len(result) < len(shortest) {
				shortest = result
			}
		}
	}

	return shortest
}

// Stack of ints
type Stack []rune

// NewStack makes a new Stack.
func NewStack() *Stack {
	var s []rune
	return (*Stack)(&s)
}

// Pop removes the topmost value on the Stack.
func (s *Stack) Pop() rune {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

// Push adds a value to the top of the Stack.
func (s *Stack) Push(i rune) {
	*s = append(*s, i)
}

// Peek retrieves the topmost value on the Stack without removing it.
func (s *Stack) Peek() rune {
	v := (*s)[len(*s)-1]
	return v
}

// Size returns the number of elements on the Stack.
func (s *Stack) Size() int {
	return len(*s)
}

// String returns the Stack as a string, from bottom to top.
func (s *Stack) String() string {
	return string(*s)
}
