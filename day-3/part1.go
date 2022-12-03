package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type void struct{}

func keysFromMap(m map[rune]struct{}) []rune {
	keys := make([]rune, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func getDuplicatedItems(dep1 []string, dep2 []string) []rune {
	set := make(map[rune]struct{})
	var v void

	for _, el := range dep1 {
		for _, el2 := range dep2 {
			if el == el2 {
				char := []rune(el)[0]
				set[char] = v
			}
		}
	}

	return keysFromMap(set)
}

func getPriority(item rune) int {
	if item < 'a' {
		return int(item) - 38
	}
	return int(item) - 96
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	line := 1
	for fileScanner.Scan() {
		val := fileScanner.Text()
		dep1 := strings.Split(val[:len(val)/2], "")
		dep2 := strings.Split(val[len(val)/2:], "")
		duplicates := getDuplicatedItems(dep1, dep2)
		fmt.Println(line, duplicates)
		for _, el := range duplicates {
			sum += getPriority(el)
		}

		line++

	}
	fmt.Println("Sum of duplicated items priorities: " + fmt.Sprint(sum))
}
