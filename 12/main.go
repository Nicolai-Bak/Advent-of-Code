package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const ROUNDS int = 10_000

var shortest int = math.MaxInt64

var fileLines []string

func main() {
	readFile("./data.txt")

	var grid [][]rune

	//Create grid
	for _, line := range fileLines {
		var newLine []rune
		for _, char := range line {
			newLine = append(newLine, char)
		}
		grid = append(grid, newLine)
	}

	//Find S and call recursive method
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'S' {
				recursiveSearch(grid, i, j, 0)
			}
		}
	}
	fmt.Println(shortest)
}

func recursiveSearch(grid [][]rune, i int, j int, currentLength int) {
	curr := grid[i][j]
	if curr == 'S' {
		curr = 'a'
	}

	if curr == 'E' {
		if currentLength < shortest {
			shortest = currentLength
			return
		}
	}

	grid[i][j] = '.'

	for _, g := range grid {
		fmt.Printf("%c\n", g)
	}
	fmt.Println()

	if i < len(grid)-1 {
		if grid[i+1][j] != '.' && (curr >= grid[i+1][j]-1 || grid[i+1][j] == 'E') {
			recursiveSearch(grid, i+1, j, currentLength+1)
		}
	}
	if j < len(grid[i])-1 {
		if grid[i][j+1] != '.' && (curr >= grid[i][j+1]-1 || grid[i][j+1] == 'E') {
			recursiveSearch(grid, i, j+1, currentLength+1)
		}
	}
	if i > 0 {
		if grid[i-1][j] != '.' && (curr >= grid[i-1][j]-1 || grid[i-1][j] == 'E') {
			recursiveSearch(grid, i-1, j, currentLength+1)
		}
	}
	if j > 0 {
		if grid[i][j-1] != '.' && (curr >= grid[i][j-1]-1 || grid[i][j-1] == 'E') {
			recursiveSearch(grid, i, j-1, currentLength+1)
		}
	}
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
