package day9

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solution() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	points := []Point{}
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(text[0])
		y, _ := strconv.Atoi(text[1])
		points = append(points, Point{x: x, y: y})
	}

	return Part1(points)
}

type Point struct {
	x int
	y int
}

func Part1(points []Point) int {
	// for _, p := range points {
	// 	println(p.x, p.y)
	// }
	maxArea := 0
	// maxPa := Point{x: -1, y: -1}
	// maxPb := Point{x: -1, y: -1}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			pa := points[i]
			pb := points[j]

			dx := pa.x - pb.x
			if dx < 0 {
				dx *= -1
			}
			dx += 1

			dy := pa.y - pb.y
			if dy < 0 {
				dy *= -1
			}
			dy += 1

			area := dx * dy

			if maxArea < area {
				maxArea = area
				// maxPa = pa
				// maxPb = pb
			}
		}
	}
	return maxArea
}

type Rect struct {
	edges []Point // diagonal edges
	area  int
}

func findRanges(points []Point) {
	colRanges := map[int][]int{}
	rowRanges := map[int][]int{}

	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		p2 := points[i+1]

		if p1.x == p2.x {
			val, ok := rowRanges[p1.x]
			if ok {
				minVal := val[0]
				maxVal := val[1]

				if p1.y < minVal {
					minVal = p1.y
				}
				if p2.y < minVal {
					minVal = p2.y
				}

				if maxVal < p1.y {
					maxVal = p1.y
				}
				if maxVal < p2.y {
					maxVal = p2.y
				}

				rowRanges[p1.x] = []int{minVal, maxVal}
			} else {
				minVal := 0
				if p1.y < p2.y {
					minVal = p1.y
				} else {
					minVal = p2.y
				}
				maxVal := 0
				if minVal == p1.y {
					maxVal = p2.y
				} else {
					maxVal = p1.y
				}
				rowRanges[p1.x] = []int{minVal, maxVal}
			}
		} else if p1.y == p2.y {
			val, ok := colRanges[p1.y]
			if ok {
				minVal := val[0]
				maxVal := val[1]

				if p1.x < minVal {
					minVal = p1.x
				}
				if p2.x < minVal {
					minVal = p2.x
				}

				if maxVal < p1.x {
					maxVal = p1.x
				}
				if maxVal < p2.x {
					maxVal = p2.x
				}
				colRanges[p1.y] = []int{minVal, maxVal}
			}
		} else {
			panic("Diagonal transition is not allowed in input")
		}
	}
}
