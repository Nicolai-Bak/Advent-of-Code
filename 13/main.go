package main

import (
	"bufio"
	"fmt"
	"os"
)

type Value struct {
	value  int
	values []Value
}

type Line struct {
	values []Value
}

type Pair struct {
	lines [2]Line
}

var fileLines []string
var pairs []Pair

func main() {
	readFile("./data.txt")

	for i := 0; i < len(fileLines); i += 3 {
		line1 := fileLines[i][1 : len(fileLines[i])-1]
		fmt.Println(line1)
	}

}

// func CreateLine(s string) Line {
// 	for i := 0; i < len(s); i++ {
// 		if s[i] == '[' {
// 			CreateValue(s[i:])
// 		}
// 	}
// }

// func CreateValue(s string) Value {

// }

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
