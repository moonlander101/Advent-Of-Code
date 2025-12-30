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

	return Part2(points)
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

func checkIfInEdge(p Point, edge []Point) bool {
	p1 := edge[0]
	p2 := edge[1]
	if p1.x == p2.x {
		if p.x != p1.x {
			return false
		}
		m := p1.y
		mi := p2.y
		if m < mi {
			m = p2.y
			mi = p1.y
		}

		if p.y >= mi && p.y <= m {
			return true
		}
	} else {
		if p.y != p1.y {
			return false
		}
		m := p1.x
		mi := p2.x
		if m < mi {
			m = p2.x
			mi = p1.x
		}

		if p.x >= mi && p.x <= m {
			return true
		}
	}
	return false
}

func checkIfInside(p Point, edges [][]Point) bool {
	// println("Checking point", p.x, p.y)
	cuts := 0
	for _, e := range edges {
		p1 := e[0]
		p2 := e[1]
		// println("Edge from", p1.x, p1.y, "to", p2.x, p2.y)

		if checkIfInEdge(p, e) {
			// println("Point on the edge, so Valid")
			return true
		}

		if p1.x == p2.x {
			// println("Vertical edge cannot intersect")
			continue
		} else {
			// println("Horizontal edge")
			m := p1.x
			mi := p2.x
			if m < mi {
				m = p2.x
				mi = p1.x
			}
			if p.x >= m || p.x < mi {
				// println("Not aligned with horizontal edge (", m, mi, ")", ", Ray at x: ", p.x)
				continue
			}
			if p1.y < p.y {
				// println("Ray crosses horizontal edge")
				cuts += 1
			}
		}
	}
	// println("Total cuts:", cuts, "Inside:", cuts%2 != 0)
	return cuts%2 != 0
}

func rectHasBoundaryInside(pa, pb Point, edges [][]Point) bool {
	minX, maxX := pa.x, pb.x
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	minY, maxY := pa.y, pb.y
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	for _, e := range edges {
		a, b := e[0], e[1]

		if a.x == b.x {
			// vertical edge x = a.x, y in [y1,y2]
			x := a.x
			y1, y2 := a.y, b.y
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			// if it lies on rectangle border, allow it
			if x == minX || x == maxX {
				continue
			}

			// does it go through rectangle interior
			if x > minX && x < maxX {
				// overlap with (minY,maxY)
				if y2 > minY && y1 < maxY {
					return true
				}
			}
		} else {
			// horizontal edge y = a.y, x in [x1,x2]
			y := a.y
			x1, x2 := a.x, b.x
			if x1 > x2 {
				x1, x2 = x2, x1
			}

			// if it lies on rectangle border, allow it
			if y == minY || y == maxY {
				continue
			}

			// does it go through rectangle interior
			if y > minY && y < maxY {
				// overlap with (minX,maxX)
				if x2 > minX && x1 < maxX {
					return true
				}
			}
		}
	}

	return false
}

func Part2(points []Point) int {
	edges := [][]Point{}
	l := len(points)
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%l]
		if p1.x == p2.x {
			// println("Horizontal line")
			edges = append(edges, []Point{p1, p2})
		} else if p1.y == p2.y {
			// println("Vertical line")
			edges = append(edges, []Point{p1, p2})
		} else {
			panic("Diagonal line not supported")
		}
	}

	maxArea := 0
	for i := range points {
		// println("===========")
		for j := i + 1; j < len(points); j++ {
			pa := points[i]
			pb := points[j]
			// println("Consider pa: ", pa.x, pa.y)
			// println("Consider pb: ", pb.x, pb.y)
			pc := Point{x: pa.x, y: pb.y}
			pd := Point{x: pb.x, y: pa.y}

			cInArea := checkIfInside(pc, edges)
			dInArea := checkIfInside(pd, edges)
			println(cInArea, dInArea)
			if !cInArea || !dInArea {
				// println("Invalid rect")
				continue
			}

			if rectHasBoundaryInside(pa, pb, edges) {
				continue
			}

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
			}
		}
	}

	return maxArea
}
