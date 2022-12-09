package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const size = 800

var grid [size][size]int

var fileLines []string

var counter int

var Hx = size / 2
var Hy = Hx
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
			prevHy := Hy
			prevHx := Hx
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

			if Hy > Ty+1 || Hy < Ty-1 || Hx > Tx+1 || Hx < Tx-1 {
				grid[prevHy][prevHx] = 1
				Ty = prevHy
				Tx = prevHx
			}
		}
	}

	for _, line := range grid {
		for _, c := range line {
			counter += c
		}
		// fmt.Println(line)
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
