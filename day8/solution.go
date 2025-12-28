package day8

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

func equal(a Point, b Point) bool {
	return a.x == b.x && a.y == b.y && a.z == b.z
}

func includes(a Point, arr []Point) bool {
	for _, p := range arr {
		if equal(a, p) {
			return true
		}
	}
	return false
}

func mergeWithoutDuplucates(arrA []Point, arrB []Point) []Point {
	merged := arrA
	for _, pb := range arrB {
		if includes(pb, merged) {
			continue
		}
		merged = append(merged, pb)
	}
	return merged
}

func Solution() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	points := []Point{}
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		z, _ := strconv.Atoi(nums[2])
		points = append(points, Point{x: x, y: y, z: z})
	}
	return Part2(points)
}

func calcDistance(a Point, b Point) float64 {
	dx := math.Abs(float64(a.x - b.x))
	dy := math.Abs(float64(a.y - b.y))
	dz := math.Abs(float64(a.z - b.z))
	d := math.Sqrt(dx*dx + dy*dy + dz*dz)
	return d
}

type Pair struct {
	a int
	b int
	d float64
}

func Part1(points []Point) int {
	distances := []Pair{}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			distance := calcDistance(points[i], points[j])
			added := false
			if !added {
				distances = append(distances, Pair{a: i, b: j, d: distance})
			}
		}
	}

	slices.SortFunc(distances, func(a Pair, b Pair) int {
		return cmp.Compare(a.d, b.d)
	})
	println("Distances sorted with", len(distances))
	for i := range distances {
		pair := distances[i]
		pa := points[pair.a]
		pb := points[pair.b]
		fmt.Printf("%d. Point Pair (%d,%d,%d) - (%d,%d,%d): %.2f\n", i, pa.x, pa.y, pa.z, pb.x, pb.y, pb.z, pair.d)
	}

	components := [][]Point{}
	// MAGIC VALUE for distance slice range based on what they have asked in example and actual input.
	// idk how to pick this conditionally
	for i, pair := range distances[:1000] {
		pa := points[pair.a]
		pb := points[pair.b]
		if i == 0 {
			comp := []Point{}
			comp = append(comp, pa, pb)
			components = append(components, comp)
			continue
		}
		compWithPa := -1
		compWithPb := -1
		for idx, comp := range components {
			includesPa := includes(pa, comp)
			includesPb := includes(pb, comp)

			if includesPa && includesPb {
				compWithPa = idx
				compWithPb = idx
			} else if includesPa {
				compWithPa = idx
			} else if includesPb {
				compWithPb = idx
			}
		}

		if compWithPa == -1 && compWithPb == -1 {
			comp := []Point{}
			comp = append(comp, pa, pb)
			components = append(components, comp)
			continue
		} else if compWithPa == -1 && compWithPb != -1 {
			components[compWithPb] = append(components[compWithPb], pa)
		} else if compWithPa != -1 && compWithPb == -1 {
			components[compWithPa] = append(components[compWithPa], pb)
		} else {
			compA := components[compWithPa]
			compB := components[compWithPb]
			if compWithPa == compWithPb {
				continue
			}
			merged := mergeWithoutDuplucates(compA, compB)
			finalComponents := [][]Point{}
			for idx, c := range components {
				if idx != compWithPa && idx != compWithPb {
					finalComponents = append(finalComponents, c)
				}
			}

			components = finalComponents
			components = append(components, merged)
		}
	}

	fmt.Println("Components:")
	for i, comp := range components {
		fmt.Printf("Component %d: %d:", i+1, len(comp))
		for _, p := range comp {
			fmt.Printf("(%d,%d,%d) ", p.x, p.y, p.z)
		}
		fmt.Println()
	}

	multipliers := []int{}
	for _, comp := range components {
		inc := false
		l := len(comp)
		if slices.Contains(multipliers, l) {
			inc = true
		}
		if !inc {
			multipliers = append(multipliers, l)
		}
	}

	slices.Sort(multipliers)
	total := 0
	l := len(multipliers)
	for i := l - 1; i >= max(l-3, 0); i-- {
		if total == 0 {
			total = multipliers[i]
			continue
		}
		total *= multipliers[i]
	}

	return total
}

func Part2(points []Point) int {
	distances := []Pair{}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			distance := calcDistance(points[i], points[j])
			added := false
			if !added {
				distances = append(distances, Pair{a: i, b: j, d: distance})
			}
		}
	}
	slices.SortFunc(distances, func(a Pair, b Pair) int {
		return cmp.Compare(a.d, b.d)
	})

	components := [][]Point{}
	connectedBoxes := 0
	lastPair := Pair{}
	for i, pair := range distances {
		pa := points[pair.a]
		pb := points[pair.b]
		if i == 0 {
			comp := []Point{}
			comp = append(comp, pa, pb)
			components = append(components, comp)
			connectedBoxes += 2
			continue
		}
		compWithPa := -1
		compWithPb := -1
		for idx, comp := range components {
			includesPa := includes(pa, comp)
			includesPb := includes(pb, comp)

			if includesPa && includesPb {
				compWithPa = idx
				compWithPb = idx
			} else if includesPa {
				compWithPa = idx
			} else if includesPb {
				compWithPb = idx
			}
		}

		if compWithPa == -1 && compWithPb == -1 {
			comp := []Point{}
			comp = append(comp, pa, pb)
			components = append(components, comp)
			connectedBoxes += 2
			continue
		} else if compWithPa == -1 && compWithPb != -1 {
			components[compWithPb] = append(components[compWithPb], pa)
			connectedBoxes += 1
		} else if compWithPa != -1 && compWithPb == -1 {
			components[compWithPa] = append(components[compWithPa], pb)
			connectedBoxes += 1
		} else {
			compA := components[compWithPa]
			compB := components[compWithPb]
			if compWithPa == compWithPb {
				continue
			}
			merged := mergeWithoutDuplucates(compA, compB)
			finalComponents := [][]Point{}
			for idx, c := range components {
				if idx != compWithPa && idx != compWithPb {
					finalComponents = append(finalComponents, c)
				}
			}

			components = finalComponents
			components = append(components, merged)
		}

		if connectedBoxes == len(points) {
			lastPair = pair
			break
		}
	}
	pa := points[lastPair.a]
	pb := points[lastPair.b]
	// fmt.Printf("Last connection for full circuit (%d, %d, %d) (%d, %d, %d)\n", pa.x, pa.y, pa.z, pb.x, pb.y, pb.z)
	return pa.x * pb.x
}
