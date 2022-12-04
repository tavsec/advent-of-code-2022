package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitRange(id string) (int, int) {
	splited := strings.Split(id, "-")
	parsed1, _ := strconv.ParseInt(splited[0], 0, 32)
	parsed2, _ := strconv.ParseInt(splited[1], 0, 32)
	return int(parsed1), int(parsed2)
}

func overlaps(id1 ID, id2 ID) bool {
	return (id1.End >= id2.Start && id1.Start <= id2.Start) || (id2.End >= id1.Start && id2.Start <= id1.Start)
}

type ID struct {
	Start int
	End   int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	counter := 0
	for fileScanner.Scan() {
		val := fileScanner.Text()
		ids := strings.Split(val, ",")

		start, end := splitRange(ids[0])
		id1 := ID{Start: start, End: end}
		start, end = splitRange(ids[1])
		id2 := ID{Start: start, End: end}
		if overlaps(id1, id2) {
			counter++
		}
	}

	fmt.Println("Number of IDs that overlaps: " + fmt.Sprint(counter))
}
