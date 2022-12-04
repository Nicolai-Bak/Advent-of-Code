package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "strings"
)

func main() {

	filePath := os.Args[1]
	readFile, _ := os.Open(filePath)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var currentSum rune

	for _, line := range fileLines {

		pairs := strings.Split(line, ",")

		pair0 := strings.Split(pairs[0], "-")
		pair1 := strings.Split(pairs[1], "-")

		pair00, _ := strconv.Atoi(pair0[0])
		pair01, _ := strconv.Atoi(pair0[1])
		pair10, _ := strconv.Atoi(pair1[0])
		pair11, _ := strconv.Atoi(pair1[1])

		// if pair00 <= pair10 && pair01 >= pair11 {
		// 	currentSum++
		// 	continue
		// }

		// if pair10 <= pair00 && pair11 >= pair01 {
		// 	currentSum++
		// 	continue
		// }

		if pair00 <= pair10 && pair01 >= pair10 {
			currentSum++
			continue
		}

		if pair00 >= pair10 && pair00 <= pair11 {
			currentSum++
			continue
		}
	}

	fmt.Println(currentSum)
}
