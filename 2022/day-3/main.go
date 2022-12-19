package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Printf("%v", LowerPriorities())
}

func Part1() string {
	input, err := os.Open("demo-input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		// -1 to compensate for empty substring which returns 1+ length of string
		num := strings.Count(line, "") - 1
		var1 := line[0 : num/2]
		var2 := line[num/2 : num]

		compare := strings.Compare(var1, var2)

		fmt.Printf("%v\n", compare)

	}

	return "test"
}

// Utils

func LowerPriorities() map[string]string {
	priorities := map[string]string{}

	for r := 'a'; r <= 'z'; r++ {

	}

	return priorities
}

// 1. Find the number of characters in each line of input
// 2. Split it in half and assign each half to a variable
// 3. Check if anything in var1 matches in var2
// 4. Create a map of characters a-z with priorities of 1-26
// 5. Create a map of characters A-Z with priorities of 27-52
// 6. Compare matched results from step 3 against the map to find the priority number for each character
// 7. Find the sum of all of the resulting priority numbers
