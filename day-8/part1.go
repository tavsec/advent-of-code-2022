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

	visibleTrees := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			if canBeVisible(trees[i][j], i, j, trees) {
				visibleTrees++
			}
		}
	}

	fmt.Println("Number of visible trees: ", visibleTrees)
}

func canBeVisible(tree, i, j int, trees [][]int) bool {
	higherOrEqualTrees := 0
	for m := i - 1; m >= 0; m-- {
		if trees[m][j] >= tree {
			higherOrEqualTrees++
		}
	}

	if higherOrEqualTrees == 0 {
		return true
	}

	higherOrEqualTrees = 0
	for m := i + 1; m < len(trees[i]); m++ {
		if trees[m][j] >= tree {
			higherOrEqualTrees++
		}
	}

	if higherOrEqualTrees == 0 {
		return true
	}

	higherOrEqualTrees = 0
	for m := j - 1; m >= 0; m-- {
		if trees[i][m] >= tree {
			higherOrEqualTrees++
		}
	}

	if higherOrEqualTrees == 0 {
		return true
	}

	higherOrEqualTrees = 0
	for m := j + 1; m < len(trees); m++ {
		if trees[i][m] >= tree {
			higherOrEqualTrees++
		}
	}

	if higherOrEqualTrees == 0 {
		return true
	}

	return false
}
