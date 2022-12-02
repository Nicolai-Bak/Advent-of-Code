package main

import (
	"bufio"
	"fmt"
	"os"
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

	const draw = 3
	const win = 6
	// const X = 1
	// const Y = 2
	// const Z = 3
	const A = 1
	const B = 2
	const C = 3

	var currentSum int

	for _, line := range fileLines {

		switch line {
		case "A X":
			currentSum += C
		case "B X":
			currentSum += A
		case "C X":
			currentSum += B
		case "A Y":
			currentSum += draw + A
		case "B Y":
			currentSum += draw + B
		case "C Y":
			currentSum += draw + C
		case "A Z":
			currentSum += win + B
		case "B Z":
			currentSum += win + C
		case "C Z":
			currentSum += win + A
		}

	}
	// for _, line := range fileLines {

	// 	switch line {
	// 	case "A X":
	// 		currentSum += draw + X
	// 	case "B X":
	// 		currentSum += X
	// 	case "C X":
	// 		currentSum += win + X
	// 	case "A Y":
	// 		currentSum += win + Y
	// 	case "B Y":
	// 		currentSum += draw + Y
	// 	case "C Y":
	// 		currentSum += Y
	// 	case "A Z":
	// 		currentSum += Z
	// 	case "B Z":
	// 		currentSum += win + Z
	// 	case "C Z":
	// 		currentSum += draw + Z
	// 	}

	// }

	fmt.Println(currentSum)
}
