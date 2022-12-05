package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	// containers := [][]rune{
	// 	{'Z', 'N'},
	// 	{'M', 'C', 'D'},
	// 	{'P'},
	// }
	containers := [][]rune{
		{'S', 'T', 'H', 'F', 'W', 'R'},
		{'S', 'G', 'D', 'Q', 'W'},
		{'B', 'T', 'W'},
		{'D', 'R', 'W', 'T', 'N', 'Q', 'Z', 'J'},
		{'F', 'B', 'H', 'G', 'L', 'V', 'T', 'Z'},
		{'L', 'P', 'T', 'C', 'V', 'B', 'S', 'G'},
		{'Z', 'B', 'R', 'T', 'W', 'G', 'P'},
		{'N', 'G', 'M', 'T', 'C', 'J', 'R'},
		{'L', 'G', 'B', 'W'},
	}

	// fmt.Printf("%c\n", containers)

	for _, line := range fileLines {

		splitted := strings.Split(line, " ")
		amount, _ := strconv.Atoi(splitted[1])
		from, _ := strconv.Atoi(splitted[3])
		to, _ := strconv.Atoi(splitted[5])

		// for i := 0; i < amount; i++ {
		// 	lenFrom := len(containers[from-1])

		// 	containers[to-1] = append(containers[to-1], containers[from-1][lenFrom-1])
		// 	containers[from-1] = containers[from-1][0 : lenFrom-1]
		// }

		lenFrom := len(containers[from-1])

		containers[to-1] = append(containers[to-1], containers[from-1][lenFrom-amount:lenFrom]...)

		containers[from-1] = containers[from-1][0 : lenFrom-amount]
		// fmt.Printf("%c\n", containers)
	}

	fmt.Printf("%c\n", containers)

	fmt.Println(currentSum)
}
