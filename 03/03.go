package main

import (
	"bufio"
	"fmt"
	"os"
)

// Coordinate represents an offset of x inches from left edge and y inches
// from top edge.
type Coordinate struct {
	x int
	y int
}

// Claim represents a rectangle area on the fabric.
type Claim struct {
	id int
	c  Coordinate
	w  int // width
	h  int // height
}

func main() {
	input, err := getInput("03.input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grid := BuildClaimGrid(input)

	overlaps := CountClaimGridOverlaps(grid)
	fmt.Println(overlaps)

	viableId := FindViableClaim(grid)
	fmt.Println(viableId)
}

func getInput(filename string) ([]Claim, error) {
	var claims []Claim

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := Claim{}
		coordinate := Coordinate{}
		fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &c.id, &coordinate.x, &coordinate.y, &c.w, &c.h)
		// this needs to be after the scan, since it's by value?
		c.c = coordinate
		claims = append(claims, c)
	}

	return claims, nil
}

// BuildClaimGrid takes a list of Elf Claims and returns a map[] with key
// representing offset from left edge and offset from top edge,
// and array of claim IDs for that square inch.
func BuildClaimGrid(claims []Claim) map[Coordinate][]int {
	grid := make(map[Coordinate][]int)
	for _, claim := range claims {
		for x := 0; x < claim.w; x++ {
			for y := 0; y < claim.h; y++ {
				c := Coordinate{x: claim.c.x + x, y: claim.c.y + y}
				grid[c] = append(grid[c], claim.id)
			}
		}
	}

	return grid
}

// CountClaimGridOverlaps counts the number of coordinates that have
// 2 or more claims recorded.
func CountClaimGridOverlaps(grid map[Coordinate][]int) int {
	overlaps := 0
	for _, v := range grid {
		if len(v) > 1 {
			overlaps++
		}
	}
	return overlaps
}

// FindViableClaim takes a claims grid and returns a claim ID which has no
// overlaps, and -1 if none are found.
func FindViableClaim(grid map[Coordinate][]int) int {
	claimExists := make(map[int]bool)
	claimOverlaps := make(map[int]int)

	// count the number of times each claim has overlapped with other claims.
	// overlapping with multiple claims in same square will be counted multiple
	// times, but that's fine, as long as 0 means no overlaps
	for _, claims := range grid {
		if len(claims) == 1 {
			// still need to record that this claim id existed
			claimExists[claims[0]] = true
		} else if len(claims) > 1 {
			for _, id := range claims {
				claimExists[id] = true
				claimOverlaps[id]++
			}
		}
	}

	for id := range claimExists {
		if claimOverlaps[id] == 0 {
			return id
		}
	}

	return -1
}
