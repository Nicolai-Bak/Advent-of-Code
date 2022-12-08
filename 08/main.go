package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileLines []string
var counter int = 0

func main() {
	readFile("./data.txt")

	height := len(fileLines)
	width := len(fileLines[0])

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			fmt.Println("it")
			counter += count(i, j)
		}
	}

	counter += height*2 + (width-2)*2
	fmt.Println(counter)
}

func count(i int, j int) int {
	curr := fileLines[i][j]
	if isVisibleVer(fileLines[i][:j], curr) {
		fmt.Printf("%c\n", curr)
		return 1
	}
	if isVisibleVer(fileLines[i][j+1:], curr) {
		fmt.Printf("%c\n", curr)
		return 1
	}
	if isVisibleHor(fileLines[:i], j, curr) {
		fmt.Printf("%c\n", curr)
		return 1
	}
	if isVisibleHor(fileLines[i+1:], j, curr) {
		fmt.Printf("%c\n", curr)
		return 1
	}
	return 0
}

func isVisibleVer(arr string, curr byte) bool {
	for i := 0; i < len(arr); i++ {
		if curr <= arr[i] {
			return false
		}
	}
	return true
}

func isVisibleHor(arr []string, index int, curr byte) bool {
	for i := 0; i < len(arr); i++ {
		if curr <= arr[i][index] {
			return false
		}
	}
	return true
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
