package day4

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Solution() int {
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	grid := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, text)
	}

	res := Part2(grid)
	return res
}

func Part1(grid []string) int {
	total := 0
	for i := range len(grid) {
		for j := range len(grid[0]) {
			count := 0
			cell := grid[i][j]
			if string(cell) != "@" {
				continue
			}
			// println("Considering cell ", i, j)
			for _, p := range []int{-1, 0, 1} {
				if i+p >= len(grid) || i+p < 0 {
					continue
				}
				for _, k := range []int{-1, 0, 1} {
					if p == 0 && k == 0 {
						continue
					}
					if j+k >= len(grid[0]) || j+k < 0 {
						continue
					}
					if string(grid[i+p][j+k]) == "@" {
						// println("Grid", i+p, j+k, "has @")
						count += 1
					}
				}
			}
			if count < 4 {
				println("Assigning on ", i, j)
				total += 1
			}
		}
	}
	return total
}

func Part2(grid []string) int {
	total := 0
	removedAny := true
	for removedAny {
		// println("=========================================")
		removingCells := [][]int{}
		for i := range len(grid) {
			for j := range len(grid[0]) {
				count := 0
				cell := grid[i][j]
				if string(cell) != "@" {
					continue
				}
				// println("Considering cell ", i, j)
				for _, p := range []int{-1, 0, 1} {
					if i+p >= len(grid) || i+p < 0 {
						continue
					}
					for _, k := range []int{-1, 0, 1} {
						if p == 0 && k == 0 {
							continue
						}
						if j+k >= len(grid[0]) || j+k < 0 {
							continue
						}
						if string(grid[i+p][j+k]) == "@" {
							// println("Grid", i+p, j+k, "has @")
							count += 1
						}
					}
				}
				if count < 4 {
					// println("Assigning on ", i, j)
					total += 1
					removingCells = append(removingCells, []int{i, j})
				}
			}
		}
		if len(removingCells) == 0 {
			removedAny = false
			continue
		}

		for _, cell := range removingCells {
			i := cell[0]
			j := cell[1]

			row := []byte(grid[i])
			row[j] = '.'
			grid[i] = string(row)
		}
		// for _, row := range grid {
		// 	println(row)
		// }
	}
	return total
}
