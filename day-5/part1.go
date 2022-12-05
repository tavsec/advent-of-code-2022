package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Ship struct {
	Crates map[int][]rune
}

func (s *Ship) addNewLineOfCrates(line string) {
	for index, el := range line {
		if el <= 'Z' && el >= 'A' {
			stackIndex := getStackIndexFromCharIndex(index)
			_, keyExists := s.Crates[stackIndex]
			if keyExists {
				s.Crates[stackIndex] = append(s.Crates[stackIndex], el)
			} else {
				s.Crates[stackIndex] = []rune{el}

			}
		}
	}
}

func (s *Ship) reversAllCrates() {
	for index, stack := range s.Crates {
		s.Crates[index] = reverseArray(stack)
	}
}

func (s *Ship) executeInstruction(count, from, to int) {
	for i := 0; i < count; i++ {
		lastElIndex := len(s.Crates[from]) - 1
		el := s.Crates[from][lastElIndex]
		s.Crates[to] = append(s.Crates[to], el)
		s.Crates[from] = s.Crates[from][:lastElIndex]
	}
}

func (s *Ship) print() {
	for index := 0; index < len(s.Crates); index++ {
		print(index + 1)
		print(" ")
		for _, create := range s.Crates[index] {
			print("[" + string(create) + "]")
			print(" ")
		}
		println()
	}
}

func (s *Ship) printTopCrates() {
	for index := 0; index < len(s.Crates); index++ {
		el := s.Crates[index][len(s.Crates[index])-1]
		print(string(el))
	}
}

func getStackIndexFromCharIndex(indexInString int) int {
	return int((0.25)*float32(indexInString)+(0.75)) - 1
}

func reverseArray[V rune | float32](originalArray []V) []V {
	for i, j := 0, len(originalArray)-1; i < j; i, j = i+1, j-1 {
		originalArray[i], originalArray[j] = originalArray[j], originalArray[i]
	}
	return originalArray
}

func parseInstruction(line string) (int, int, int) {
	r := regexp.MustCompile(`\d+`)
	result := r.FindAllStringSubmatch(line, -1)
	num1, _ := strconv.Atoi(result[0][0])
	num2, _ := strconv.Atoi(result[1][0])
	num3, _ := strconv.Atoi(result[2][0])
	return num1, num2, num3
}

func main() {
	ship := Ship{make(map[int][]rune)}
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	isOnInstructions := false

	for fileScanner.Scan() {
		val := fileScanner.Text()
		if val == "" {
			ship.reversAllCrates()
			isOnInstructions = true
			continue
		}

		if !isOnInstructions {
			ship.addNewLineOfCrates(val)
			continue
		}

		counter, from, to := parseInstruction(val)
		ship.executeInstruction(counter, from-1, to-1)
	}

	ship.print()
	ship.printTopCrates()

}
