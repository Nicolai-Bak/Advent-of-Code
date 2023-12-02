package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value     rune
	friends   [4]*Node
	usedSteps int
}

const rows = 41
const columns = 136

var startingNode *Node

var fileLines []string

func main() {
	readFile("./data.txt")

	var nodes [rows][]Node

	for i := 0; i < rows; i++ {
		nodes[i] = make([]Node, columns)
	}

	//Create grid
	for row, line := range fileLines {
		for column, char := range line {
			nodes[row][column].value = char
			nodes[row][column].usedSteps = 900

			if row > 0 {
				nodes[row][column].friends[0] = &nodes[row-1][column]
			}
			if row < rows-1 {
				nodes[row][column].friends[1] = &nodes[row+1][column]
			}
			if column > 0 {
				nodes[row][column].friends[2] = &nodes[row][column-1]
			}
			if column < columns-1 {
				nodes[row][column].friends[3] = &nodes[row][column+1]
			}

			if char == 'S' {
				nodes[row][column].value = 'a'

				startingNode = &nodes[row][column]
			}
		}
	}

	recursiveSearch(startingNode, -1)
}

func recursiveSearch(node *Node, usedSteps int) {
	node.usedSteps = usedSteps + 1

	if node.value == 'a' {
		node.usedSteps = 0
	} else if node.value == 'E' {
		fmt.Println(node.usedSteps)
		return
	}

	for i := 0; i < 4; i++ {
		if node.friends[i] != nil {
			if node.usedSteps+1 < node.friends[i].usedSteps &&
				((node.friends[i].value == 'E' && node.value+1 >= 'z') ||
					(node.friends[i].value != 'E' && node.value+1 >= node.friends[i].value)) {
				recursiveSearch(node.friends[i], node.usedSteps)
			}
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
