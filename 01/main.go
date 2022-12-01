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

	var currentCalSum, maxOne, maxTwo, maxThree int

	fileLines = append(fileLines, "")

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

	var one = fmt.Sprintf("1: %d", maxOne)
	fmt.Println(one)
	var two = fmt.Sprintf("2: %d", maxTwo)
	fmt.Println(two)
	var three = fmt.Sprintf("3: %d", maxThree)
	fmt.Println(three)

	var max = fmt.Sprintf("sum: %d", maxOne+maxTwo+maxThree)
	fmt.Println(max)

}
