package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Solution() int {
	file, err := os.Open("input.txt")

	check(err)

	scanner := bufio.NewScanner(file)

	var total int = 0

	for scanner.Scan() {
		text := scanner.Text()
		checkRange := strings.Split(text, "-")
		start, err := strconv.Atoi(checkRange[0])
		check(err)

		end, err := strconv.Atoi(checkRange[1])
		check(err)

		for i := start; i <= end; i++ {
			has := Part2(strconv.Itoa(i))
			if has {
				total += i
			}

		}

	}

	return total
}

func Part1(s string) bool {
	l := len(s)
	if l%2 != 0 {
		return false
	}
	i := l / 2
	s1 := s[0:i]
	s2 := s[i:l]

	if s1 == s2 {
		return true
	}
	return false
}

func Part2(s string) bool {
	l := len(s)

	for i := 1; i <= (l / 2); i++ {
		isEqual := true
		next := i
		for j := i; j+i <= l; j += i {
			// println(s[0:i], s[j:j+i])
			if s[0:i] != s[j:j+i] {
				isEqual = false
			}
			next += i
		}
		if l-next < i && l-next != 0 {
			continue
		}
		if isEqual {
			// println("Evaluated as equal", s)
			return true
		}
	}

	return false
}
