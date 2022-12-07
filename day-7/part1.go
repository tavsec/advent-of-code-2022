package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sumOfDirsUnder100k = int64(0)

type File struct {
	Name string
	Size int64
}

type Directory struct {
	Parent      *Directory
	Name        string
	Directories []*Directory
	Files       []File
}

func (d *Directory) pushFile(file File) {
	d.Files = append(d.Files, file)
}

func (d *Directory) pushDirectory(directory *Directory) {
	d.Directories = append(d.Directories, directory)
}

func (d *Directory) getSubDirByName(name string) *Directory {
	for _, dir := range d.Directories {
		if dir.Name == name {
			return dir
		}
	}
	return nil
}

func (d *Directory) getSize() int64 {
	sum := int64(0)
	for _, file := range d.Files {
		sum += file.Size
	}

	for _, directory := range d.Directories {
		dirSize := directory.getSize()

		sum += dirSize
	}

	if sum <= 100000 {
		sumOfDirsUnder100k += sum
	}

	return sum
}

func parseCommand(line string) []string {
	args := strings.Split(line, " ")
	return args[1:]
}

func main() {
	system := Directory{Parent: nil, Name: "/"}
	currentDir := &system
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Couldn't open input.txt")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		val := fileScanner.Text()
		if strings.HasPrefix(val, "$") {
			args := parseCommand(val)
			if args[0] == "cd" {
				if args[1] == "/" {
					currentDir = &system
				} else {
					if args[1] == ".." {
						currentDir = currentDir.Parent
					} else {
						currentDir = currentDir.getSubDirByName(args[1])
					}
				}
			}
		} else {
			if strings.HasPrefix(val, "dir") {
				info := strings.Split(val, " ")
				dir := Directory{Parent: currentDir, Name: info[1]}
				currentDir.pushDirectory(&dir)

			} else {
				info := strings.Split(val, " ")
				size, _ := strconv.ParseInt(info[0], 10, 64)
				file := File{Name: info[1], Size: size}
				currentDir.pushFile(file)

			}
		}

	}

	system.getSize()
	fmt.Println("Sum of directory size under 100k:")
	fmt.Println(sumOfDirsUnder100k)
}
