package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const knots = 10

var grid [800][800]int

var fileLines []string

var counter int

var Tx [knots]int
var Ty [knots]int

func main() {
	readFile("./data.txt")

	for i := 0; i < len(Ty); i++ {
		Ty[i] = 400
		Tx[i] = 400
	}
	grid[Ty[knots-1]][Tx[knots-1]] = 1

	for _, line := range fileLines {
		cmd := strings.Split(line, " ")
		dir := cmd[0]
		steps, _ := strconv.Atoi(cmd[1])

		for step := 0; step < steps; step++ {
			prevTx := Tx
			prevTy := Ty
			if dir == "D" {
				Ty[0]++
			}
			if dir == "U" {
				Ty[0]--
			}
			if dir == "L" {
				Tx[0]--
			}
			if dir == "R" {
				Tx[0]++
			}

			if Ty[0] > Ty[1]+1 || Ty[0] < Ty[1]-1 || Tx[0] > Tx[1]+1 || Tx[0] < Tx[1]-1 {
				Ty[1] = prevTy[0]
				Tx[1] = prevTx[0]
			}

			for j := 2; j < knots; j++ {
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
						if Ty[j-1] == prevTy[j] {
							Ty[j] = Ty[j-1]
						}
						if Tx[j-1] == prevTx[j] {
							Tx[j] = Tx[j-1]
						}
					}
				}
			}

			grid[Ty[knots-1]][Tx[knots-1]] = 1

		}
	}

	for _, line := range grid {
		for _, c := range line {
			counter += c
		}
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
