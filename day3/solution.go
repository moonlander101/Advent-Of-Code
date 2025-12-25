package day3

import (
	"bufio"
	"os"
	"strconv"
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

	total := 0

	for scanner.Scan() {
		text := scanner.Text()
		total += Part2(text)
	}
	return total
}

func Part1(s string) int {
	first := -1
	idx1 := -1
	second := -2
	idx2 := -1
	for i, c := range s {
		num, err := strconv.Atoi(string(c))
		check(err)
		if num > first {
			second = first
			idx2 = idx1
			first = num
			idx1 = i
		}
	}
	l := len(s)
	if idx1 > idx2 {
		if idx1 == l-1 {
			res, err := strconv.Atoi(string(s[idx2]) + string(s[idx1]))
			check(err)
			// println(res)
			return res
		}
		second = -1
		idx2 = -1
		for i := idx1 + 1; i < l; i++ {
			num, err := strconv.Atoi(string(s[i]))
			check(err)
			if num > second {
				second = num
				idx2 = i
			}
		}
	}
	res, err := strconv.Atoi(string(s[idx1]) + string(s[idx2]))
	check(err)
	// println(res)
	return res
}

func compareNumericStrings(a string, b string) bool {
	if len(a) != len(b) {
		panic("Not equal")
	}

	for i := 0; i < len(a); i++ {
		d1, _ := strconv.Atoi(a[i : i+1])
		d2, _ := strconv.Atoi(b[i : i+1])
		if d1 < d2 {
			return false
		}
	}
	return true
}

func Part2(s string) int {
	for range len(s) - 12 {
		maxWhenRemoved := 0
		maxVal := s[1:]
		// println("=====")
		for j := 1; j < len(s)-1; j++ {
			val := s[0:j] + s[j+1:]
			if compareNumericStrings(val, maxVal) {
				maxWhenRemoved = j
				maxVal = val
			}
		}

		// final check
		val := s[:len(s)-1]
		if compareNumericStrings(val, maxVal) {
			maxWhenRemoved = len(s) - 1
		}

		// println("Removing ", maxWhenRemoved)
		s = s[0:maxWhenRemoved] + s[maxWhenRemoved+1:]
		// println("final", s)

	}
	res, _ := strconv.Atoi(s)
	return res
}
