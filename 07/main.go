package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const needed int64 = 30000000

var totalDisk int64 = 70000000

var smallestDir int64 = totalDisk
var total int64

var missing int64

var fileLines []string
var dirSizes []int64

func main() {
	readFile("./data.txt")

	for _, line := range fileLines {
		if line[:1] != "$" && line[:3] != "dir" {
			stringSize := strings.Split(line, " ")[0]
			size, _ := strconv.ParseInt(stringSize, 10, 64)
			totalDisk -= size
		}
	}
	missing = needed - totalDisk

	recursive(0)

	for _, size := range dirSizes {
		total += size
	}

	fmt.Println(total)
	fmt.Println(smallestDir)
}

func recursive(startIndex int) (int, int64) {
	var currSize int64

	for i := startIndex; i < len(fileLines); i++ {
		if fileLines[i] == "$ cd .." {
			if currSize >= missing && currSize < smallestDir {
				smallestDir = currSize
			} else if currSize <= 100000 {
				dirSizes = append(dirSizes, currSize)
			}
			return i, currSize
		}
		if strings.Contains(fileLines[i], "$ cd ") {
			newI, size := recursive(i + 1)
			currSize += size
			i = newI
		}
		if fileLines[i][:1] != "$" && fileLines[i][:3] != "dir" {
			stringSize := strings.Split(fileLines[i], " ")[0]
			size, _ := strconv.ParseInt(stringSize, 10, 64)
			currSize += size
		}
	}
	if currSize >= missing && currSize < smallestDir {
		smallestDir = currSize
	} else if currSize <= 100000 {
		dirSizes = append(dirSizes, currSize)
	}
	return len(fileLines) - 1, currSize
}

func readFile(path string) {
	filePath := path
	readFile, _ := os.Open(filePath)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
}
