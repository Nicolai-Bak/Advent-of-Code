package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var grid [800][800]int

// var path [22][26]string

var fileLines []string

var counter int

var Hx = 400
var Hy = 400

var Tx [9]int
var Ty [9]int

func main() {
	readFile("./data.txt")

	for i := 0; i < len(Ty); i++ {
		Ty[i] = Hy
		Tx[i] = Hx
	}
	grid[Ty[8]][Tx[8]] = 1

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

			prevTx := Tx
			prevTy := Ty

			if Hy > Ty[0]+1 || Hy < Ty[0]-1 || Hx > Tx[0]+1 || Hx < Tx[0]-1 {
				Ty[0] = prevHy
				Tx[0] = prevHx
			}

			for j := 1; j < 9; j++ {
				if Ty[j-1] > Ty[j]+1 || Ty[j-1] < Ty[j]-1 || Tx[j-1] > Tx[j]+1 || Tx[j-1] < Tx[j]-1 {
					Ty[j] = prevTy[j-1]
					Tx[j] = prevTx[j-1]

					if prevTy[j-1] != Ty[j-1] && prevTx[j-1] != Tx[j-1] {
						if prevTy[j-1] == prevTy[j] {
							Ty[j] = Ty[j-1]
						}
						if prevTx[j-1] == prevTx[j] {
							Tx[j] = Tx[j-1]
						}
					}
					// if j == 8 {
					// 	continue
					// }
					// if Ty[j-1] == Ty[j+1] {
					// 	Ty[j] = Ty[j-1]
					// }
					// if Tx[j-1] == Tx[j+1] {
					// 	Tx[j] = Tx[j-1]
					// }
				}
			}

			grid[Ty[8]][Tx[8]] = 1

		}
		// for k := 0; k < len(path); k++ {
		// 	for l := 0; l < len(path[k]); l++ {
		// 		path[k][l] = "."
		// 	}
		// }
		// for j := 8; j >= 0; j-- {
		// 	path[Ty[j]][Tx[j]] = fmt.Sprintf("%d", j+1)
		// }
		// path[Hy][Hx] = "H"
		// for _, line := range path {
		// 	fmt.Println(line)
		// }
		// fmt.Println()
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
