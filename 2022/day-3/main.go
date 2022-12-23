package main

import (
	"bufio"
	"fmt"
	"io"
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
	i := 0
	group := make([]string, 3)
	scanner := bufio.NewReader(Input())
	for {
		line, _, err := scanner.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		group[i] = string(line)
		i++
		if i == 3 {
			Badge(group)
			i = 0
			continue
		}

	}
	return 0
}

// Utils

func Input() *os.File {
	input, err := os.Open("demo-input.txt")
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

func Badge(group []string) int {
	i := 0
	j := 1
	s := []string{}
	for _, v := range group[i] {
		if strings.Contains(group[j], string(v)) {
			s = append(s, string(v))
		}
	}

	for _, val := range s {
		int := Priorities()[string(val)]

	}
	return 0
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
