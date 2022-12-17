package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var score int = 0

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "X") {
			score += 1
		} else if strings.Contains(scanner.Text(), "Y") {
			score += 2
		} else if strings.Contains(scanner.Text(), "Z") {
			score += 3
		}

		switch line := scanner.Text(); line {
		case "A Y", "B Z", "C X":
			score += 6
		case "A X", "B Y", "C Z":
			score += 3
		}
	}
	fmt.Println(score)
}
