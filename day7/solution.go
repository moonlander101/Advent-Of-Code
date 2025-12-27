package day7

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

	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return Part2(input)
}

func splitTachyon(a int, arr []int) ([]int, bool) {
	leftAdded := false
	rightAdded := false
	tachIdx := -1
	for i, v := range arr {
		if v == a-1 {
			leftAdded = true
		}
		if v == a+1 {
			rightAdded = true
		}
		if v == a {
			tachIdx = i
		}
	}
	if tachIdx == -1 {
		return arr, false
	}
	arr = append(arr[:tachIdx], arr[tachIdx+1:]...)
	if !leftAdded {
		arr = append(arr, a-1)
	}
	if !rightAdded {
		arr = append(arr, a+1)
	}
	return arr, true
}

func Part1(input []string) int {
	tacPositions := []int{}
	splitCount := 0
	for i, c := range input[0] {
		if c == 'S' {
			tacPositions = append(tacPositions, i)
		}
	}
	for i, row := range input {
		if i == 0 {
			continue
		}
		for j, c := range row {
			if c == '^' {
				didSplit := false
				tacPositions, didSplit = splitTachyon(j, tacPositions)
				if didSplit {
					splitCount += 1
				}
			}
		}
	}
	return splitCount
}

type Timeline struct {
	idx       int
	timelines int
}

func splitTimeline(idx int, arr []Timeline) []Timeline {
	toBeSplit := arr[0]
	arrIdx := 0
	assigned := false
	for i, t := range arr {
		if t.idx == idx {
			assigned = true
			toBeSplit = t
			arrIdx = i
		}
	}
	if !assigned {
		return arr
	}
	arr = append(arr[:arrIdx], arr[arrIdx+1:]...)

	left := -1
	right := -1
	for j, t := range arr {
		if t.idx == toBeSplit.idx-1 {
			// println("Found leftside")
			left = j
		}

		if t.idx == toBeSplit.idx+1 {
			// println("Found rightside")
			right = j
		}
	}

	if left != -1 {
		arr[left].timelines += toBeSplit.timelines
	} else {
		newTimeline := Timeline{idx: toBeSplit.idx - 1, timelines: toBeSplit.timelines}
		arr = append(arr, newTimeline)
	}

	if right != -1 {
		arr[right].timelines += toBeSplit.timelines
	} else {
		newTimeline := Timeline{idx: toBeSplit.idx + 1, timelines: toBeSplit.timelines}
		arr = append(arr, newTimeline)
	}

	return arr
}

func Part2(input []string) int {
	tacTimelines := []Timeline{}
	for i, c := range input[0] {
		if c == 'S' {
			// println("initial at ", i)
			t := Timeline{idx: i, timelines: 1}
			tacTimelines = append(tacTimelines, t)
			break
		}
	}
	for i, row := range input {
		// println("---")
		if i == 0 {
			continue
		}
		for j, c := range row {
			if c == '^' {
				// println("splitting at ", j)
				tacTimelines = splitTimeline(j, tacTimelines)
				// for _, t := range tacTimelines {
				// 	println(t.idx, t.timelines)
				// }
			}
		}
	}

	total := 0
	for _, t := range tacTimelines {
		total += t.timelines
	}
	return total
}
