package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type OrderedStack struct {
	topArray []int64
}

func (orderedStack *OrderedStack) AddNewNumber(sum int64) {
	for index, el := range orderedStack.topArray {
		if sum > el {
			orderedStack.MoveDownElement(index, sum)
			return
		}
	}
}

func (orderedStack *OrderedStack) MoveDownElement(index int, value int64) {
	if index >= len(orderedStack.topArray) {
		return
	}
	tmpVal := orderedStack.topArray[index]
	orderedStack.topArray[index] = value
	orderedStack.MoveDownElement(index+1, tmpVal)
}

func (orderedStack *OrderedStack) GetTotalSum() int64 {
	var sum int64 = 0
	for _, el := range orderedStack.topArray {
		sum += el
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	orderedStack := OrderedStack{[]int64{0, 0, 0}}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var sum int64 = 0
	for fileScanner.Scan() {
		val := fileScanner.Text()

		if len(val) == 0 {
			orderedStack.AddNewNumber(sum)
			sum = 0
			continue

		}

		number, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			panic("Couldn't parse " + val + " to int64.")
		}
		sum += number
	}

	fmt.Println("1. max calories: " + fmt.Sprint(orderedStack.topArray[0]))
	fmt.Println("2. max calories: " + fmt.Sprint(orderedStack.topArray[1]))
	fmt.Println("3. max calories: " + fmt.Sprint(orderedStack.topArray[2]))
	fmt.Println("Top 3 calories sum: " + fmt.Sprint(orderedStack.GetTotalSum()))
}
