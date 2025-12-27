package day6

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solution() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	numbers := [][]int{}
	operations := []string{}

	// Part2 is easier to parse the input as is
	grid := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, text)
		split := strings.Fields(text)
		arr := []int{}
		opline := false
		for _, c := range split {
			n, err := strconv.Atoi(c)
			if err != nil {
				opline = true
				operations = append(operations, c)
			} else {
				arr = append(arr, n)
			}
		}
		if !opline {
			numbers = append(numbers, arr)
		}
	}

	// return Part1(numbers, operations)
	return Part2(grid)
}

func Part1(numbers [][]int, operations []string) int {
	total := 0
	for j := range numbers[0] {
		res := 0
		op := operations[j]
		for i := range numbers {
			if res == 0 {
				res = numbers[i][j]
				continue
			}
			switch op {
			case "*":
				res *= numbers[i][j]
			case "+":
				res += numbers[i][j]
			}
		}
		total += res
	}
	return total
}

func Part2(input []string) int {
	currentSet := 0
	nextSet := 1
	total := 0
	for currentSet != len(input[0])-1 {
		for input[len(input)-1][nextSet] == ' ' && nextSet < len(input[0])-1 {
			nextSet += 1
		}
		res := 0
		op := string(input[len(input)-1][currentSet])

		endOfCurSet := nextSet - 2
		if nextSet == len(input[0])-1 {
			endOfCurSet = nextSet
		}
		for i := endOfCurSet; i >= currentSet; i-- {
			numStr := ""
			for j := 0; j < len(input)-1; j++ {
				numStr += string(input[j][i])
			}
			numStr = strings.TrimSpace(numStr)
			num, _ := strconv.Atoi(numStr)
			// println(num)
			if res == 0 {
				res = num
				continue
			}
			switch op {
			case "+":
				res += num
			case "*":
				res *= num
			}
		}
		// println(res)
		total += res
		currentSet = nextSet
		nextSet += 1
	}
	return total
}
