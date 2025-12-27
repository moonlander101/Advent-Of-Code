package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	ranges := []string{} // "a-b"
	ids := []int{}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		ranges = append(ranges, text)
	}

	for scanner.Scan() {
		text := scanner.Text()
		n, _ := strconv.Atoi(text)
		ids = append(ids, n)
	}

	return Part3(ranges, ids)
}

func checkIfNumInRange(num int, r string) bool {
	min, _ := strconv.Atoi(r[:strings.Index(r, "-")])
	max, _ := strconv.Atoi(r[strings.Index(r, "-")+1:])
	if num >= min && num <= max {
		return true
	}
	return false
}

func Part1(ranges []string, ids []int) int {
	total := 0
	for _, id := range ids {
		for _, r := range ranges {
			found := checkIfNumInRange(id, r)
			if found {
				total += 1
				break
			}
		}
	}
	return total
}

func includes(a int, arr []int) bool {
	for _, i := range arr {
		if i == a {
			return true
		}
	}
	return false
}

// TODO: Merge and check difference for eacg range
func mergeRanges(ranges []string) []string {
	onlyOutside := false
	for !onlyOutside {
		onlyOutside = true
		finalRanges := []string{}
		for _, r := range ranges {
			// println("===========")
			minR, _ := strconv.Atoi(r[:strings.Index(r, "-")])
			maxR, _ := strconv.Atoi(r[strings.Index(r, "-")+1:])
			// println("max min: ", maxR, minR)
			merged := false
			for idx, fr := range finalRanges {
				minFr, _ := strconv.Atoi(fr[:strings.Index(fr, "-")])
				maxFr, _ := strconv.Atoi(fr[strings.Index(fr, "-")+1:])
				// println("max min f ", maxFr, minFr)
				if maxR <= maxFr && minR >= minFr {
					merged = true
					onlyOutside = false
					// println("Inside this range")
					continue
				}
				if minR > maxFr || maxR < minFr {
					// println("Outside the range")
					continue
				}
				if maxR >= maxFr && minR >= minFr {
					merged = true
					onlyOutside = false
					// println("extend to right")
					finalRanges[idx] = fmt.Sprintf("%d-%d", minFr, maxR)
					continue
				}
				if maxR <= maxFr && minR <= minFr {
					// println("extend to left")
					onlyOutside = false
					merged = true
					finalRanges[idx] = fmt.Sprintf("%d-%d", minR, maxFr)
					continue
				}
				if maxR >= maxFr && minR <= minFr {
					// println("extend both ways")
					onlyOutside = false
					merged = true
					finalRanges[idx] = fmt.Sprintf("%d-%d", minR, maxR)
					break
				}
			}
			if !merged {
				// println("Adding as seperate interval")
				finalRanges = append(finalRanges, r)
			}
		}
		ranges = finalRanges
	}

	return ranges
}

func Part2(ranges []string, ids []int) int {
	freshItems := []int{}
	for _, r := range ranges {
		start, _ := strconv.Atoi(r[:strings.Index(r, "-")])
		end, _ := strconv.Atoi(r[strings.Index(r, "-")+1:])
		for i := start; i <= end; i++ {
			if includes(i, freshItems) {
				continue
			}
			freshItems = append(freshItems, i)
		}
	}
	return len(freshItems)
}

func Part3(ranges []string, ids []int) int {
	finalRanges := mergeRanges(ranges)
	total := 0
	for _, r := range finalRanges {
		minFr, _ := strconv.Atoi(r[:strings.Index(r, "-")])
		maxFr, _ := strconv.Atoi(r[strings.Index(r, "-")+1:])

		total += maxFr - minFr + 1
	}
	return total
}
