package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	allSpace        = int64(70000000)
	neededFreeSpace = int64(30000000)
)

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

	return sum
}

func (d *Directory) isBiggerOrEqual(size int64) bool {
	return d.getSize() >= size
}

func (d *Directory) getArrayOfBiggerOrEqualSubdirectories(size int64, existingArray []Directory) []Directory {
	arr := make([]Directory, 0)
	if d.isBiggerOrEqual(size) {
		arr = append(arr, *d)
	}

	for _, directory := range d.Directories {
		eligible := directory.getArrayOfBiggerOrEqualSubdirectories(size, make([]Directory, 0))
		arr = append(arr, eligible...)
	}

	return append(existingArray, arr...)
}

func parseCommand(line string) []string {
	args := strings.Split(line, " ")
	return args[1:]
}

func getTheSmallestDir(dirs []Directory) Directory {
	minVal := int64(math.MaxInt64)
	minDir := Directory{Files: make([]File, 0)}
	for _, dir := range dirs {
		if dir.getSize() < minVal {
			minDir = dir
			minVal = dir.getSize()
		}
	}
	return minDir
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

	spaceNeeded := neededFreeSpace - (allSpace - system.getSize())
	eligebleDirectories := system.getArrayOfBiggerOrEqualSubdirectories(spaceNeeded, make([]Directory, 0))
	minDir := getTheSmallestDir(eligebleDirectories)
	fmt.Println(minDir.getSize())
}
