package main

import (
	"bufio"
	"fmt"
	"os"
)

// Coordinate represents an offset of x units from left edge and y units
// from top edge.
type Coordinate struct {
	x int
	y int
}

// Location represents a specific sector of the grid.
type Location struct {
	closest int
	dist    int
	tied    bool
}

func main() {
	points, err := getInput("06.input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	area := BruteforceLargestArea(points)
	fmt.Println(area)
}

func getInput(filename string) ([]Coordinate, error) {
	coordinates := []Coordinate{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := Coordinate{}
		fmt.Sscanf(scanner.Text(), "%d, %d", &c.x, &c.y)
		coordinates = append(coordinates, c)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return coordinates, nil
}

// FindBoundaries takes a list of coordinates, and returns two coordinates
// containing the lowest x/y values seen and the highest x/y values seen
// respectively.
func FindBoundaries(coordinates []Coordinate) (Coordinate, Coordinate) {
	minX := -1
	maxX := -1
	minY := -1
	maxY := -1
	for _, c := range coordinates {
		if minX == -1 || c.x < minX {
			minX = c.x
		}
		if maxX == -1 || c.x > maxX {
			maxX = c.x
		}
		if minY == -1 || c.y < minY {
			minY = c.y
		}
		if maxY == -1 || c.y > maxY {
			maxY = c.y
		}
	}
	return Coordinate{x: minX, y: minY}, Coordinate{x: maxX, y: maxY}
}

// Dist takes two coordinates and returns the Manhattan distance between them.
func Dist(a Coordinate, b Coordinate) int {
	x := a.x - b.x
	if x < 0 {
		x = -x
	}
	y := a.y - b.y
	if y < 0 {
		y = -y
	}
	return x + y
}

// BruteforceLargestArea takes a list of coordinates, and returns the size of
// the largest non-infinite area that is closest to some coordinate.
func BruteforceLargestArea(coordinates []Coordinate) int {
	lower, upper := FindBoundaries(coordinates)

	// marker for whether some coordinate is known to have infinite area
	infinite := make(map[int]bool)
	grid := make(map[Coordinate]int)

	for x := lower.x; x <= upper.x; x++ {
		for y := lower.y; y <= upper.y; y++ {
			minDist := -1
			minC := -1
			for id, c := range coordinates {
				d := Dist(Coordinate{x: x, y: y}, c)
				if minDist == -1 || d < minDist {
					minDist = d
					minC = id
				}
			}
			// if this is one of the boundaries, mark this area as infinite
			if x == lower.x || x == upper.x || y == lower.y || y == upper.y {
				infinite[minC] = true
			}
			grid[Coordinate{x: x, y: y}] = minC
		}
	}

	areas := make(map[int]int)
	for _, id := range grid {
		if !infinite[id] {
			areas[id]++
		}
	}
	maxArea := 0
	for _, area := range areas {
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
