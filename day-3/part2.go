package main

import (
	"bufio"
	"fmt"
	"os"
)

type Group struct {
	items [][]rune
}

func (g *Group) addItems(items []rune) {
	g.items = append(g.items, items)
}

func (g *Group) getBadge() rune {
	firstItemSet := g.items[:1]
	otherItems := g.items[1:]
	counter := 0
	for _, item := range firstItemSet[0] {
		for _, otherItem := range otherItems {
			if hasRune(otherItem, item) {
				counter++
			}
		}
		if counter == len(otherItems) {
			return item
		}

		counter = 0
	}

	return 0
}

func hasRune(a []rune, c rune) bool {
	for _, e := range a {
		if e == c {
			return true
		}
	}
	return false
}

func getPriority(item rune) int {
	if item < 'a' {
		return int(item) - 38
	}
	return int(item) - 96
}

func main() {

	var groups []Group

	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	index := 0
	groupsIndex := 0
	for fileScanner.Scan() {
		val := fileScanner.Text()

		if index%3 == 0 {
			groups = append(groups, Group{})
			groupsIndex++
		}

		groups[groupsIndex-1].addItems([]rune(val))
		index++

	}

	sum := 0
	for _, group := range groups {

		badge := group.getBadge()
		sum += getPriority(badge)
	}

	fmt.Println("Sum of groups badge's priority: " + fmt.Sprint(sum))

}
