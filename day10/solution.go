package day10

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInput(line string) ([]byte, [][]int, []int) {
	target := []byte{}
	buttons := [][]int{}
	joltage := []int{}
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c == ' ' {
			continue
		}
		if c == '[' {
			count := 0
			for j := i + 1; line[j] != ']'; j += 1 {
				target = append(target, line[j])
				count += 1
			}
			i += count
			continue
		}

		if c == '(' {
			count := 0
			button := []int{}
			for j := i + 1; line[j] != ')'; j += 1 {
				if line[j] == ',' {
					continue
				}
				num, err := strconv.Atoi(string(line[j]))
				if err != nil {
					panic("Couldnt convert number")
				}
				button = append(button, num)
				count += 1
			}
			i += count
			buttons = append(buttons, button)
		}

		if c == '{' {
			start := i + 1
			end := i + 1
			for j := i + 1; line[j] != '}'; j += 1 {
				end = j
			}
			parts := strings.Split(line[start:end+1], ",")
			for _, p := range parts {
				n, err := strconv.Atoi(strings.TrimSpace(p))
				if err != nil {
					panic("Bruh")
				}
				joltage = append(joltage, n)
			}

		}
	}

	// fmt.Printf("buttons: %v\n", buttons)
	// fmt.Printf("target: %s\n", string(target))
	// fmt.Printf("joltage: %v\n", joltage)
	return target, buttons, joltage
}

func Solution() int {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		target, buttons, _ := parseInput(text)
		tar, butt, ss := convToBinaryRep(target, buttons)
		total += Part1B(tar, butt, ss)
	}
	return total
}

func powInt(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func convToBinaryRep(target []byte, buttons [][]int) (int, []int, int) {
	tar := 0
	l := len(target)
	pow := 0
	for i := l - 1; i >= 0; i-- {
		if i == l-1 {
			pow = 1
		} else {
			pow *= 2
		}
		if target[i] == '#' {
			tar += pow
		}
	}
	butts := []int{}
	for _, but := range buttons {
		val := 0
		for _, i := range but {
			pow := l - i - 1
			val += powInt(2, pow)
		}
		butts = append(butts, val)
	}

	return tar, butts, powInt(2, len(target))
}

type Node struct {
	visited  bool
	distance int
}

func findMinimumUnvisitedNode(visited []Node) int {
	k := 0
	d := 100000000000000
	for i, n := range visited {
		if d > n.distance && n.visited == false {
			d = n.distance
			k = i
		}
	}
	return k
}

func Part1B(target int, buttons []int, stateSize int) int {
	// println("target: ", target)
	// fmt.Printf("buttons %v\n", buttons)
	visited := []Node{}
	for i := range stateSize {
		if i == 0 {
			visited = append(visited, Node{visited: false, distance: 0})
			continue
		}
		visited = append(visited, Node{visited: false, distance: 1000000000000})
	}
	for visited[target].visited == false {
		// fmt.Printf("====\nCurrenmt visited %v\n", visited)
		n := findMinimumUnvisitedNode(visited)
		visited[n].visited = true
		// println("Now at ", n)
		for _, b := range buttons {
			nextNode := n ^ b
			// println("Pressing", b, "Next Node", nextNode)
			if visited[nextNode].visited {
				continue
			} else {
				oldDistance := visited[nextNode].distance
				newDistance := visited[n].distance + 1
				// println("Old, ", oldDistance, "New, ", newDistance)
				if min(oldDistance, newDistance) == newDistance {
					visited[nextNode].distance = newDistance
				}
			}
		}
	}

	return visited[target].distance
}
