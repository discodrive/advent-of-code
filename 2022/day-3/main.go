package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("%v", Part2())
}

func Part1() int {
	var score int = 0

	scanner := bufio.NewScanner(Input())

	for scanner.Scan() {
		line := scanner.Text()
		// -1 to compensate for empty substring which returns 1+ length of string
		num := strings.Count(line, "") - 1
		var1 := line[0 : num/2]
		var2 := line[num/2 : num]

		score += Sum(Score(var1, var2))

	}
	return score
}

func Part2() int {
	score := 0
	counter := 0
	i := 0
	scanner := bufio.NewScanner(Input())

	for scanner.Scan() {
		if counter != 3 {
			fmt.Println(i)
			counter++
		} else {
			counter = 0
		}
		i++
	}

	return score
}

// Utils

func Input() *os.File {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func Priorities() map[string]int {
	priorities := make(map[string]int)
	i := 1
	t := 27
	for r := 'a'; r <= 'z'; r++ {
		priorities[string(r)] = i
		i++
		R := unicode.ToUpper(r)
		priorities[string(R)] = t
		t++
	}
	return priorities
}

func Score(var1 string, var2 string) []int {
	i := []int{}
	for _, v := range var1 {
		if strings.Contains(var2, string(v)) {
			int := Priorities()[string(v)]
			i = append(i, int)
		}
	}
	return Unique(i)
}

func Unique(arr []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for e := range arr {
		if occurred[arr[e]] != true {
			occurred[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}

func Sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// 1. Find the number of characters in each line of input
// 2. Split it in half and assign each half to a variable
// 3. Check if anything in var1 matches in var2
// 4. Create a map of characters a-z with priorities of 1-26
// 5. Create a map of characters A-Z with priorities of 27-52
// 6. Compare matched results from step 3 against the map to find the priority number for each character
// 7. Find the sum of all of the resulting priority numbers
