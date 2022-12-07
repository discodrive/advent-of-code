package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%v\n", Part1())
	fmt.Printf("%v", Part2())
}

func Part1() int {
	max := 0
	sum := 0

	for _, calory := range ReadFile() {
		if calory != "" {
			cals, _ := strconv.Atoi(calory)
			sum += cals
			if sum > max {
				max = sum
			}
		} else {
			sum = 0
		}
	}
	return max
}

func Part2() int {
	elves := []int{}
	sum := 0

	for _, calory := range ReadFile() {
		if calory != "" {
			cals, _ := strconv.Atoi(calory)
			sum += cals
		} else {
			elves = append(elves, sum)
			sum = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	top3 := elves[0:3]

	return Sum(top3)
}

// UTILS

func ReadFile() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}

func Sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
