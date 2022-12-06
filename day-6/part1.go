package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	MaxItems int
	items    []rune
}

func (s *Stack) addItem(item rune) {
	lastIndex := len(s.items)
	if lastIndex < s.MaxItems {
		s.items = append(s.items, item)
		return
	}

	s.pushItem(s.MaxItems-1, item)
	s.items[s.MaxItems-1] = item
}

func (s *Stack) pushItem(index int, item rune) {
	if index < 0 {
		return
	}

	el := s.items[index]
	s.items[index] = item
	s.pushItem(index-1, el)
}

func (s *Stack) areItemsUnique() bool {
	if len(s.items) < s.MaxItems {
		return false
	}
	set := make(map[rune]struct{})
	for _, el := range s.items {
		set[el] = struct{}{}
	}

	return len(set) == s.MaxItems
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		stack := Stack{MaxItems: 4}
		val := fileScanner.Text()
		counter := 0
		for _, char := range val {
			stack.addItem(char)
			counter++
			if stack.areItemsUnique() {
				fmt.Println(string(stack.items))
				println(counter)
				return
			}
		}
	}
}
