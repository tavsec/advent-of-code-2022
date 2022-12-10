package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Register struct {
	X       int64
	Counter int64
	sum     int64
}

func (r *Register) increaseCounter() {
	r.Counter++
	if r.Counter == 20 || r.Counter == 60 || r.Counter == 100 || r.Counter == 140 || r.Counter == 180 || r.Counter == 220 {
		fmt.Println(r.X*r.Counter, r.Counter, r.X)
		r.sum += r.X * r.Counter
	}
}

func (r *Register) addX(X int64) {
	r.X += X
}

func main() {
	register := Register{X: 1, Counter: 0, sum: 0}
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		val := fileScanner.Text()
		if val == "noop" {
			register.increaseCounter()
			continue
		}
		args := strings.Split(val, " ")
		number, _ := strconv.ParseInt(args[1], 10, 64)
		register.increaseCounter()
		register.increaseCounter()
		register.addX(number)

	}

	fmt.Println("Sum of strong signals", register.sum)

}
