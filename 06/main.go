package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	run()
}

func run() {
	filePath := "./data.txt"
	readFile, _ := os.Open(filePath)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	line := fileLines[0]

	// for i := 3; i < len(line); i++ {
	// 	var arr []byte

	// 	for j := 3; j >= 0; j-- {
	// 		contains := false
	// 		for k := 0; k < len(arr); k++ {
	// 			if arr[k] == line[i-j] {
	// 				contains = true
	// 			}
	// 		}
	// 		if !contains {
	// 			arr = append(arr, line[i-j])
	// 		}
	// 	}

	// 	if len(arr) == 4 {
	// 		fmt.Println(i + 1)
	// 		break
	// 	}
	// }
	for i := 13; i < len(line); i++ {
		var arr []byte

		for j := 13; j >= 0; j-- {
			contains := false
			for k := 0; k < len(arr); k++ {
				if arr[k] == line[i-j] {
					contains = true
				}
			}
			if !contains {
				arr = append(arr, line[i-j])
			}
		}

		if len(arr) == 14 {
			fmt.Println(i + 1)
			break
		}
	}

}
