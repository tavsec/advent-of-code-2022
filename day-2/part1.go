package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

const (
	Draw = 3
	Win  = 6
	Lose = 0
)

func getNumberOfPoints(opponentChoiceStr string, ourChoiceStr string) (int, int) {
	ourChoice := mapInput(ourChoiceStr)
	opponentChoice := mapInput(opponentChoiceStr)

	fmt.Println(ourChoice, opponentChoice)
	switch ourChoice - opponentChoice {
	case 0:
		return Draw + ourChoice, Draw + opponentChoice
	case -1, 2:
		return Lose + ourChoice, Win + opponentChoice
	case -2, 1:
		return Win + ourChoice, Lose + opponentChoice
	}

	return -1, -1
}

func mapInput(choice string) int {
	switch choice {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}

	return -1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		val := fileScanner.Text()
		choices := strings.Split(val, " ")
		ourPoints, opponentPoints := getNumberOfPoints(choices[0], choices[1])

		fmt.Println(ourPoints, opponentPoints)
		sum += ourPoints
	}

	fmt.Println("Sum of our points: " + fmt.Sprint(sum))
}
