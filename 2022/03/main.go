package main

import (
	"bufio"
	"fmt"
	"os"
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

	// for _, line := range fileLines {

	// 	length := len(line)
	// 	line1 := line[0 : length/2]
	// 	line2 := line[length/2 : length]

	// 	for _, ch := range line1 {

	// 		if strings.Contains(line2, fmt.Sprintf("%c", ch)) {
	// 			var value rune = ch - 64 + 26
	// 			if ch > 96 {
	// 				value = ch - 96
	// 			}
	// 			fmt.Println(value)
	// 			fmt.Printf("%c\n", ch)
	// 			currentSum += value
	// 			break
	// 		}
	// 	}
	// }

	for i := 2; i < len(fileLines); i += 3 {

		for _, ch := range fileLines[i] {

			if strings.Contains(fileLines[i-2], fmt.Sprintf("%c", ch)) {
				if strings.Contains(fileLines[i-1], fmt.Sprintf("%c", ch)) {
					var value rune = ch - 64 + 26
					if ch > 96 {
						value = ch - 96
					}
					fmt.Println(value)
					fmt.Printf("%c\n", ch)
					currentSum += value
					break
				}
			}

		}
	}

	fmt.Println(currentSum)
}
