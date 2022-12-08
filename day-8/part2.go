package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	trees := make([][]int, 0)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	counter := 0
	for fileScanner.Scan() {
		trees = append(trees, make([]int, 0))
		val := fileScanner.Text()
		for _, tree := range val {
			value, _ := strconv.Atoi(string(tree))
			trees[counter] = append(trees[counter], value)
		}
		counter++
	}

	biggestScore := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			score := getScenicScore(trees[i][j], i, j, trees)
			if score > biggestScore {
				biggestScore = score
			}
		}
	}
	fmt.Println("Biggest scenic score: ", biggestScore)

}

func getScenicScore(tree, i, j int, trees [][]int) int {
	bottom := 0
	for m := i - 1; m >= 0; m-- {
		bottom++
		if trees[m][j] >= tree {
			break
		}
	}

	top := 0
	for m := i + 1; m < len(trees[i]); m++ {
		top++
		if trees[m][j] >= tree {
			break
		}
	}

	left := 0
	for m := j - 1; m >= 0; m-- {
		left++
		if trees[i][m] >= tree {
			break
		}
	}

	right := 0
	for m := j + 1; m < len(trees); m++ {
		right++
		if trees[i][m] >= tree {
			break
		}
	}

	return left * right * top * bottom
}
