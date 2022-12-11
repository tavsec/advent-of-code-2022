package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Screen struct {
	X                  int64
	Counter            int64
	sum                int64
	Chars              []string
	currentRowAddition int64
}

func (r *Screen) increaseCounter() {
	if r.Counter+1-r.currentRowAddition >= r.X && r.Counter-1-r.currentRowAddition <= r.X {
		r.addChar("#")
	} else {
		r.addChar(".")
	}
	r.Counter++

	if (r.Counter)%40 == 0 {
		r.currentRowAddition += 40
	}
}

func (r *Screen) addX(X int64) {
	r.X += X
}

func (r *Screen) addChar(char string) {
	r.Chars = append(r.Chars, char)
}

func (r *Screen) print() {
	for i, el := range r.Chars {
		fmt.Print(el, " ")
		if (i+1)%40 == 0 {
			fmt.Println("")
		}
	}
}

func main() {
	screen := Screen{X: 1, Counter: 0}
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		val := fileScanner.Text()
		if val == "noop" {
			screen.increaseCounter()
			continue
		}
		args := strings.Split(val, " ")
		number, _ := strconv.ParseInt(args[1], 10, 64)
		screen.increaseCounter()
		screen.increaseCounter()
		screen.addX(number)

	}

	screen.print()

}
