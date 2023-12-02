package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileLines []string
var counter int

var maxSight int

func main() {
	readFile("./data.txt")

	height := len(fileLines)
	width := len(fileLines[0])

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			// fmt.Println("it")
			counter += count(i, j)
			sight := sightVer(i, j, height) * sightHor(i, j, width)
			if sight > maxSight {
				// fmt.Println(sight)
				maxSight = sight
			}
		}
	}

	counter += height*2 + (width-2)*2
	// fmt.Println(counter)
	fmt.Println(maxSight)
}

func sightVer(row int, column int, height int) int {
	var down, up int
	// fmt.Printf("curr %c\n", fileLines[row][column])
	for i := row + 1; i < height; i++ {
		down++
		// fmt.Printf("%c\n", fileLines[i][column])
		if fileLines[i][column] >= fileLines[row][column] {
			break
		}
	}
	for i := row - 1; i >= 0; i-- {
		up++
		// fmt.Printf("%c\n", fileLines[i][column])
		if fileLines[i][column] >= fileLines[row][column] {
			break
		}
	}
	// fmt.Printf("up %d\n", up)
	// fmt.Println()
	return down * up
}

func sightHor(row int, column int, width int) int {
	var right, left int
	for i := column + 1; i < width; i++ {
		right++
		if fileLines[row][i] >= fileLines[row][column] {
			break
		}
	}
	for i := column - 1; i >= 0; i-- {
		left++
		if fileLines[row][i] >= fileLines[row][column] {
			break
		}
	}
	return left * right
}

func count(i int, j int) int {
	curr := fileLines[i][j]
	if isVisibleVer(fileLines[i][:j], curr) {
		// fmt.Printf("%c\n", curr)
		return 1
	}
	if isVisibleVer(fileLines[i][j+1:], curr) {
		// fmt.Printf("%c\n", curr)
		return 1
	}
	if isVisibleHor(fileLines[:i], j, curr) {
		// fmt.Printf("%c\n", curr)
		return 1
	}
	if isVisibleHor(fileLines[i+1:], j, curr) {
		// fmt.Printf("%c\n", curr)
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
