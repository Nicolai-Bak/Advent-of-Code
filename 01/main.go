package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	fileLines = append(fileLines, "")

	var currentCalSum, maxOne, maxTwo, maxThree int

	for _, line := range fileLines {
		if line == "" {
			if currentCalSum > maxOne {
				maxThree = maxTwo
				maxTwo = maxOne
				maxOne = currentCalSum
			} else if currentCalSum > maxTwo {
				maxThree = maxTwo
				maxTwo = currentCalSum
			} else if currentCalSum > maxThree {
				maxThree = currentCalSum
			}
			currentCalSum = 0
		}

		curr, _ := strconv.Atoi(line)
		currentCalSum = currentCalSum + curr
	}

	fmt.Println(fmt.Sprintf("sum: %d", maxOne+maxTwo+maxThree))

	var one = fmt.Sprintf("1: %d", maxOne)
	fmt.Println(one)
	var two = fmt.Sprintf("2: %d", maxTwo)
	fmt.Println(two)
	var three = fmt.Sprintf("3: %d", maxThree)
	fmt.Println(three)

}
