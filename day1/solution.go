package day1

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Solution1() int {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "input.txt")
	readFile, err := os.Open(path)

	check(err)

	scanner := bufio.NewScanner(readFile)

	scanner.Split(bufio.ScanLines)
	timesAtZero := 0
	position := 50
	fmt.Println("Position at", position)
	for scanner.Scan() {
		instruction := scanner.Text()
		rotationDir := -1
		direction := string(instruction[0])
		if direction == "R" {
			rotationDir *= -1
		}
		distance, err := strconv.Atoi(instruction[1:])
		check(err)

		position += (distance * rotationDir)
		if position < 0 {
			position += 100
		}

		position = position % 100
		if position == 0 {
			timesAtZero += 1
		}
		// fmt.Println("Position at", position)
	}

	return timesAtZero
}

func Solution2() int {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "input.txt")
	readFile, err := os.Open(path)

	check(err)

	scanner := bufio.NewScanner(readFile)

	scanner.Split(bufio.ScanLines)
	timesPastZero := 0
	position := 50
	previousPosition := 50
	// fmt.Println("Position at", position)
	for scanner.Scan() {
		// fmt.Println("====================================")
		// fmt.Println("Current Position ", position)
		instruction := scanner.Text()
		rotationDir := -1
		direction := string(instruction[0])
		if direction == "R" {
			rotationDir *= -1
		}
		distance, err := strconv.Atoi(instruction[1:])

		check(err)

		fullRotations := distance / 100

		remainingDistance := distance - (fullRotations * 100)
		// fmt.Println("Remaining Distance ", remainingDistance*rotationDir)
		// fmt.Println("Full Rotations: ", fullRotations)
		timesPastZero += fullRotations

		position += (remainingDistance * rotationDir)

		if position <= 0 {
			if previousPosition != 0 {
				timesPastZero += 1
			}
			position += 100
			// fmt.Println("Passed Zero")
			if position == 100 {
				position = 0
			}
		}

		if position >= 100 {
			position -= 100
			if previousPosition != 100 {
				timesPastZero += 1
			}
			// fmt.Println("Passed Zero")
		}
		// fmt.Println("New Position", position)
		previousPosition = position
		// fmt.Println("Position at", position)
	}

	return timesPastZero
}
