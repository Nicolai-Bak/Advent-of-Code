package main

import (
	"bufio"
	"fmt"
	"os"
)

const ROUNDS uint = 1000000000000
const HEIGHT uint = 4
const WIDTH uint = 7

var fileLines []string
var rockTypes = []Rock{
	{
		coordinates: []Coordinates{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		},
		mostLeftX:  0,
		mostRightX: 3,
		height:     1,
	},
	{
		coordinates: []Coordinates{
			{1, 0},
			{0, 1},
			{1, 1},
			{2, 1},
			{1, 2},
		},
		mostLeftX:  0,
		mostRightX: 2,
		height:     3,
	},
	{
		coordinates: []Coordinates{
			{2, 0},
			{2, 1},
			{0, 2},
			{1, 2},
			{2, 2},
		},
		mostLeftX:  0,
		mostRightX: 2,
		height:     3,
	},
	{
		coordinates: []Coordinates{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
		},
		mostLeftX:  0,
		mostRightX: 0,
		height:     4,
	},
	{
		coordinates: []Coordinates{
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 1},
		},
		mostLeftX:  0,
		mostRightX: 1,
		height:     2,
	},
}

var rocks = make([]Rock, ROUNDS)

var grid [][7]int = make([][7]int, ROUNDS*HEIGHT) //[ROUNDS*HEIGHT + 4][WIDTH]int

var current_height uint = 0
var steps uint = 0

func main() {
	readFile("./data.txt")

	var jet = fileLines[0]

	for i := range rocks {
		r := rockTypes[i%len(rockTypes)]
		rocks[i].coordinates = make([]Coordinates, len(r.coordinates))
		for j := range rocks[i].coordinates {
			rocks[i].coordinates[j].x = r.coordinates[j].x
			rocks[i].coordinates[j].y = r.coordinates[j].y
		}
		rocks[i].mostLeftX = r.mostLeftX
		rocks[i].mostRightX = r.mostRightX
		rocks[i].height = r.height
	}

	for _, rock := range rocks {
		arrangeRock(&rock)

		for i := uint(0); i < current_height+4; i++ {

			keepMoving := moveRock(&rock, jet)

			steps++
			if !keepMoving {
				break
			}
		}
		setRock(rock)
		current_height = maxHeight(rock.coordinates)

	}
	// for _, g := range grid {
	// 	fmt.Println(g)
	// }
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	fmt.Println(current_height)
	// for _, b := range rocks[len(rocks)-1].coordinates {
	// 	fmt.Println(b)
	// }
}

func setRock(rock Rock) {
	for _, c := range rock.coordinates {
		grid[c.y][c.x] = 1
	}
}

func moveRock(rock *Rock, jet string) bool {
	canMove := true
	if jet[steps%uint(len(jet))] == '<' {
		if rock.mostLeftX > 0 {
			for _, c := range rock.coordinates {
				if grid[c.y][c.x-1] == 1 {
					canMove = false
					break
				}
			}
			if canMove {
				rock.mostLeftX--
				rock.mostRightX--
				for i := range rock.coordinates {
					rock.coordinates[i].x--
				}
			}
		}
	} else {
		if rock.mostRightX < WIDTH-1 {
			for _, c := range rock.coordinates {
				if grid[c.y][c.x+1] == 1 {
					canMove = false
					break
				}
			}
			if canMove {
				rock.mostLeftX++
				rock.mostRightX++
				for i := range rock.coordinates {
					rock.coordinates[i].x++
				}
			}
		}
	}
	canMove = true
	for _, c := range rock.coordinates {
		if c.y+uint(1) > uint(len(grid)-1) {
			canMove = false
			break
		}
		if grid[c.y+1][c.x] == 1 {
			canMove = false
			break
		}
	}
	if canMove {
		for i := range rock.coordinates {
			rock.coordinates[i].y++
		}
	}
	// fmt.Println("move")
	return canMove
}

func arrangeRock(rock *Rock) {
	for i, c := range rock.coordinates {
		rock.coordinates[i].x = c.x + 2
		rock.coordinates[i].y = c.y + uint(len(grid)) - rock.height - uint(3) - current_height
	}
	rock.mostLeftX = rock.mostLeftX + 2
	rock.mostRightX = rock.mostRightX + 2

}

func maxHeight(cs []Coordinates) uint {
	var eq [7]int
	for i := len(grid) - 1; i > 0; i-- {
		if grid[i] == eq {
			return uint(len(grid) - i - 1)
		}
	}
	return uint(len(grid))
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

type Rock struct {
	coordinates []Coordinates
	mostLeftX   uint
	mostRightX  uint
	height      uint
}

type Coordinates struct {
	x uint
	y uint
}
