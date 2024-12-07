package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var dbg = func() func(format string, as ...any) {
	if os.Getenv("DEBUG") == "" {
		return func(string, ...any) {}
	}
	return func(format string, as ...any) {
		fmt.Printf(format+"\n", as...)
	}
}()

type point struct {
	x int
	y int
}

type footprint string

const (
	UP        footprint = "up"
	DOWN      footprint = "down"
	LEFT      footprint = "left"
	RIGHT     footprint = "right"
	BOTH_WAYS footprint = "+"
	UNKNOWN   footprint = "unknown"
)

type guard struct {
	// ^: [-1, 0]
	dirsMap   map[string][2]int
	dirOrder  []string
	dirIdx    int
	seen      map[point]footprint
	turnCount map[point]int
}

func newGaurd(p point, dir string) *guard {
	g := &guard{
		dirsMap: map[string][2]int{
			"^": [2]int{-1, 0},
			">": [2]int{0, 1},
			"v": [2]int{1, 0},
			"<": [2]int{0, -1},
		},
		dirOrder: []string{"^", ">", "v", "<"},
		dirIdx:   0,
		seen: map[point]footprint{
			// always starts up
			p: UP,
		},
		turnCount: make(map[point]int),
	}
	g.initDirIdx(dir)
	return g
}

func (g *guard) turnRight() {
	g.dirIdx = (g.dirIdx + 1) % 4
}

func (g *guard) markTurn(p point) {
	g.seen[p] = BOTH_WAYS
}

func (g *guard) lookAhead(p point) point {
	newX, newY := p.x+g.dirsMap[g.dirOrder[g.dirIdx]][0], p.y+g.dirsMap[g.dirOrder[g.dirIdx]][1]
	return point{newX, newY}
}

func (g *guard) determineFootprint() footprint {
	//dbg("determineFootprint")
	if g.dirIdx == 0 {
		return UP
	} else if g.dirIdx == 1 {
		return RIGHT
	} else if g.dirIdx == 2 {
		//dbg("  LEFT_RIGHT")
		return DOWN
	} else if g.dirIdx == 3 {
		return LEFT
	}
	//dbg("  %s", UNKNOWN)
	log.Fatal("determine unknown")
	return UNKNOWN
}

func (g *guard) storeToSeen(p point) (footprint, error) {

	fp := g.determineFootprint()
	if footprint, ok := g.seen[p]; !ok {
		if fp == UNKNOWN {
			log.Fatalf("determineFootprint returned %s", UNKNOWN)
		}
		g.seen[p] = fp
		return fp, nil
	} else {
		if footprint == fp {
			return UNKNOWN, errors.New("repeat")
		}
		g.updateSeen(p)
		return BOTH_WAYS, nil
	}
}

func (g *guard) updateSeen(p point) {
	g.seen[p] = BOTH_WAYS
}

func (g *guard) initDirIdx(dir string) {
	idx := slices.Index(g.dirOrder, dir)
	if idx == -1 {
		log.Fatalf("dir not found. dir: %s, idx: %d", dir, idx)
	}
	g.dirIdx = idx
}

func loadData(path string) [][]string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("os.Open: %v", err)
	}

	s := bufio.NewScanner(file)

	var matrix [][]string

	for s.Scan() {

		matrix = append(matrix, strings.Split(s.Text(), ""))
	}

	return matrix

}

// AI Helpers ---
// copyMatrix creates a deep copy of a 2D slice
func copyMatrix(matrix [][]string) [][]string {
	// Create a new slice with the same dimensions
	copy := make([][]string, len(matrix))
	for i := range matrix {
		// Create a copy of each row
		copy[i] = make([]string, len(matrix[i]))
		copySlice(copy[i], matrix[i])
	}
	return copy
}

// copySlice is a helper function to copy one slice into another
func copySlice(dst, src []string) {
	copy(dst, src)
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		fmt.Println(strings.Join(row, ""))
	}
}

func debugGuard(g *guard, currentPoint point) {
	fmt.Printf("Guard Position: (%d, %d), Direction: %s\n", currentPoint.x, currentPoint.y, g.dirOrder[g.dirIdx])
}

// ------

func main() {

	ogMatrix := loadData("inputs/day06")
	var (
		ROWS = len(ogMatrix)
		COLS = len(ogMatrix[0])
	)

	var startPoint point
	var startPos [2]int
	var startDir string
	// find gaurd
	matrix := copyMatrix(ogMatrix)
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == "^" ||
				matrix[i][j] == ">" ||
				matrix[i][j] == "v" ||
				matrix[i][j] == "<" {
				fmt.Println("found guard!")
				startPoint = point{i, j}
				startPos[0] = i
				startPos[1] = j
				startDir = matrix[i][j]
			}
		}
	}

	inBounds := func(p point) bool {
		if p.x < 0 ||
			p.y < 0 ||
			p.x >= ROWS ||
			p.y >= COLS {
			return false
		}
		return true

	}

	isObstical := func(cp point) bool {
		if matrix[cp.x][cp.y] == "#" {
			return true
		}
		return false
	}

	g := newGaurd(startPoint, startDir)
	currentPoint := startPoint
	for inBounds(currentPoint) {

		newPoint := g.lookAhead(currentPoint)
		if !inBounds(newPoint) {
			//dbg("!inBounds: %v", newPoint)
			break
		}
		if isObstical(newPoint) {
			//dbg("isObstical: %v", newPoint)
			g.turnRight()
		} else {
			_, _ = g.storeToSeen(newPoint)
			//dbg("passed storeToSeen: newPoint: %v", newPoint)
			currentPoint = newPoint
		}

	}
	// part 1
	fmt.Println(len(g.seen))

	visitedPoints := g.seen
	var loopCount int
	// var dbgCount int
	for p, _ := range visitedPoints {
		dbg("with point: %v, loopCount: %d", p, loopCount)
		if p == startPoint {
			continue
		}

		matrix = copyMatrix(ogMatrix)

		g := newGaurd(startPoint, startDir)
		matrix[p.x][p.y] = "#"

		currentPoint = startPoint

		for inBounds(currentPoint) {

			if p.x == 6 && p.y == 3 {
				dbg("x: %d, y: %d", currentPoint.x, currentPoint.y)
			}

			newPoint := g.lookAhead(currentPoint)
			if !inBounds(newPoint) {
				//dbg("!inBounds: %v", newPoint)
				break
			}
			if isObstical(newPoint) {
				if p.x == 6 && p.y == 3 {
					dbg("OBSTICAL")
				}
				//dbg("isObstical: %v", newPoint)
				if val, ok := g.seen[currentPoint]; ok {
					if val == BOTH_WAYS {

						g.turnCount[currentPoint]++
						if g.turnCount[currentPoint] > 2 {
							if p.x == 6 && p.y == 3 {
								dbg("EXITING: x: %d, y: %d", currentPoint.x, currentPoint.y)
							}
							loopCount++
							break
						}
					}
				}
				g.turnRight()
				g.markTurn(currentPoint)
			} else {
				currentPoint = newPoint
			}

		}

	}

	// part 2
	fmt.Println(loopCount)

}
