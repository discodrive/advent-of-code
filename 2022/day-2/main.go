package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Printf("%v", Part1())
	fmt.Printf("%v", Part2())
}

func Part1() int {
	var score int = 0

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		// Lose cases - Points assigned for each shape chosen
		if strings.Contains(scanner.Text(), "X") {
			score += 1
		} else if strings.Contains(scanner.Text(), "Y") {
			score += 2
		} else if strings.Contains(scanner.Text(), "Z") {
			score += 3
		}

		switch line := scanner.Text(); line {
		// Winning combinations
		case "A Y", "B Z", "C X":
			score += 6
		// Draw combinations
		case "A X", "B Y", "C Z":
			score += 3
		}
	}
	return score
}

func Part2() int {
	var score int = 0

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "X") {
			// Lose = +0
			switch line {
			case "A X":
				score += 3
			case "B X":
				score += 1
			case "C X":
				score += 2
			}
		} else if strings.Contains(line, "Y") {
			// Draw = +3
			switch line {
			case "A Y":
				score += 4
			case "B Y":
				score += 5
			case "C Y":
				score += 6
			}
		} else if strings.Contains(line, "Z") {
			// Win = +6
			switch line {
			case "A Z":
				score += 8
			case "B Z":
				score += 9
			case "C Z":
				score += 7
			}
		}
	}
	return score
}
