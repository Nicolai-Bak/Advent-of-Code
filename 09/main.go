package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var grid [800][800]int

var fileLines []string

var counter int

var Hx = 400
var Hy = 400
var Tx = Hx
var Ty = Hy

func main() {
	readFile("./data.txt")

	grid[Hy][Hx] = 1

	for _, line := range fileLines {
		cmd := strings.Split(line, " ")
		dir := cmd[0]
		steps, _ := strconv.Atoi(cmd[1])

		for i := 0; i < steps; i++ {
			if dir == "D" {
				Hy++
			}
			if dir == "U" {
				Hy--
			}
			if dir == "L" {
				Hx--
			}
			if dir == "R" {
				Hx++
			}
			if (dir == "R" || dir == "L") && Hx == Tx {
				continue
			}
			if (dir == "U" || dir == "D") && Hy == Ty {
				continue
			}
			if i != steps-1 {
				grid[Hy][Hx] = 1
				Ty = Hy
				Tx = Hx
			}
		}
	}

	for _, line := range grid {
		for _, c := range line {
			counter += c
		}
		fmt.Println(line)
	}

	fmt.Println(counter)
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
