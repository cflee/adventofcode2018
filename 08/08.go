package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Node is a node on the tree.
type Node struct {
	children []Node
	metadata []int
}

func main() {
	root, err := getInput("08.input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(sumMetadata(root))
	fmt.Println(calcValue(root))
}

func getInput(filename string) (Node, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Node{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	node, err := parseNode(scanner)
	if err != nil {
		return Node{}, err
	}

	return node, nil
}

func parseNode(scanner *bufio.Scanner) (Node, error) {
	node := Node{}

	scanner.Scan()
	s := scanner.Text()
	if err := scanner.Err(); err != nil {
		return node, err
	}
	numChildren, err := strconv.Atoi(s)
	if err != nil {
		return node, err
	}

	scanner.Scan()
	s = scanner.Text()
	if err := scanner.Err(); err != nil {
		return node, err
	}
	numMetadata, err := strconv.Atoi(s)
	if err != nil {
		return node, err
	}

	for i := 0; i < numChildren; i++ {
		n, err := parseNode(scanner)
		if err != nil {
			return node, err
		}
		node.children = append(node.children, n)
	}
	for i := 0; i < numMetadata; i++ {
		scanner.Scan()
		s = scanner.Text()
		if err := scanner.Err(); err != nil {
			return node, err
		}
		m, err := strconv.Atoi(s)
		if err != nil {
			return node, err
		}
		node.metadata = append(node.metadata, m)
	}

	return node, nil
}

func sumMetadata(node Node) int {
	sum := 0
	for _, c := range node.children {
		sum += sumMetadata(c)
	}
	for _, m := range node.metadata {
		sum += m
	}
	return sum
}

func calcValue(node Node) int {
	if len(node.children) == 0 {
		sum := 0
		for _, m := range node.metadata {
			sum += m
		}
		return sum
	}
	sum := 0
	for _, m := range node.metadata {
		if m != 0 && (m-1) < len(node.children) {
			sum += calcValue(node.children[m-1])
		}
	}
	return sum
}
