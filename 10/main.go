package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fileLines []string

var cycles int
var X int = 1
var signalStrength int

var CRT [6][40]string

func main() {
	readFile("./data.txt")

	for _, line := range fileLines {
		input := strings.Split(line, " ")
		cmd := input[0]

		if cmd == "noop" {
			position := cycles % 40
			if X-1 <= position && position <= X+1 {
				CRT[cycles/40][cycles%40] = "#"
			} else {
				CRT[cycles/40][cycles%40] = "."
			}
			cycles++
			continue
		}
		if cmd == "addx" {
			for i := 0; i < 2; i++ {
				position := cycles % 40
				if X-1 <= position && position <= X+1 {
					CRT[cycles/40][cycles%40] = "#"
				} else {
					CRT[cycles/40][cycles%40] = "."
				}
				cycles++
				if cycles == 20 || (cycles-20)%40 == 0 {
					signalStrength += cycles * X

				}
			}
			value, _ := strconv.Atoi(input[1])
			X += value
			// fmt.Println(value)
		}
	}

	fmt.Println()
	for _, line := range CRT {
		fmt.Println(line)
	}
	fmt.Println()
	fmt.Println(cycles)
	fmt.Println(X)
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
