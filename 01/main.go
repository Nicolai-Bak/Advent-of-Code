package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var prev int = 1
	var counter int

	for _, line := range fileLines {
		curr, _ := strconv.Atoi(line)

		if curr > prev {
			counter += 1
		}

		fmt.Println(line)

		prev = curr
	}

	fmt.Println(fileLines)
	fmt.Println(counter)
}
