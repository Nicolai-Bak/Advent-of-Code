package main

import (
	"fmt"
)

const ROUNDS int = 10_000

// var fileLines []string

// var monkies = []Monkey{
// 	{
// 		items:          []int{79, 98},
// 		operation:      '*',
// 		operationValue: 19,
// 		divisibleBy:    23,
// 		toWhenTrue:     2,
// 		toWhenFalse:    3,
// 	},
// 	{
// 		items:          []int{54, 65, 75, 74},
// 		operation:      '+',
// 		operationValue: 6,
// 		divisibleBy:    19,
// 		toWhenTrue:     2,
// 		toWhenFalse:    0,
// 	},
// 	{
// 		items:          []int{79, 60, 97},
// 		operation:      '.',
// 		operationValue: 0,
// 		divisibleBy:    13,
// 		toWhenTrue:     1,
// 		toWhenFalse:    3,
// 	},
// 	{
// 		items:          []int{74},
// 		operation:      '+',
// 		operationValue: 3,
// 		divisibleBy:    17,
// 		toWhenTrue:     0,
// 		toWhenFalse:    1,
// 	},
// }

var monkies = []Monkey{
	{
		items:          []int{64},
		operation:      '*',
		operationValue: 7,
		divisibleBy:    13,
		toWhenTrue:     1,
		toWhenFalse:    3,
	},
	{
		items:          []int{60, 84, 84, 65},
		operation:      '+',
		operationValue: 7,
		divisibleBy:    19,
		toWhenTrue:     2,
		toWhenFalse:    7,
	},
	{
		items:          []int{52, 67, 74, 88, 51, 61},
		operation:      '*',
		operationValue: 3,
		divisibleBy:    5,
		toWhenTrue:     5,
		toWhenFalse:    7,
	},
	{
		items:          []int{67, 72},
		operation:      '+',
		operationValue: 3,
		divisibleBy:    2,
		toWhenTrue:     1,
		toWhenFalse:    2,
	},
	{
		items:          []int{80, 79, 58, 77, 68, 74, 98, 64},
		operation:      '.',
		operationValue: 0,
		divisibleBy:    17,
		toWhenTrue:     6,
		toWhenFalse:    0,
	},
	{
		items:          []int{62, 53, 61, 89, 86},
		operation:      '+',
		operationValue: 8,
		divisibleBy:    11,
		toWhenTrue:     4,
		toWhenFalse:    6,
	},
	{
		items:          []int{86, 89, 82},
		operation:      '+',
		operationValue: 2,
		divisibleBy:    7,
		toWhenTrue:     3,
		toWhenFalse:    0,
	},
	{
		items:          []int{92, 81, 70, 96, 69, 84, 83},
		operation:      '+',
		operationValue: 4,
		divisibleBy:    3,
		toWhenTrue:     4,
		toWhenFalse:    5,
	},
}

func main() {
	for round := 0; round < ROUNDS; round++ {
		for i, monkey := range monkies {
			for _, item := range monkey.items {
				monkies[i].inspections++
				if monkey.operationValue == 0 {
					item *= item
				} else if monkey.operation == '*' {
					item *= monkey.operationValue
				} else if monkey.operation == '+' {
					item += monkey.operationValue
				}

				// item /= 3
				// item %= (23 * 19 * 13 * 17)
				item %= (13 * 19 * 5 * 2 * 17 * 11 * 7 * 3)

				if item%monkey.divisibleBy == 0 {
					monkies[monkey.toWhenTrue].items = append(monkies[monkey.toWhenTrue].items, item)
				} else {
					monkies[monkey.toWhenFalse].items = append(monkies[monkey.toWhenFalse].items, item)
				}
			}
			monkies[i].items = []int{}
		}
	}

	var most int = 0
	var secondMost int = 0
	for _, monkey := range monkies {
		if monkey.inspections > most {
			secondMost = most
			most = monkey.inspections
		} else if monkey.inspections > secondMost {
			secondMost = monkey.inspections
		}
		fmt.Println(monkey.items)
		fmt.Println(monkey.inspections)
	}

	fmt.Println(most * secondMost)
}

type Monkey struct {
	inspections    int
	items          []int
	operation      rune
	operationValue int
	divisibleBy    int
	toWhenTrue     int
	toWhenFalse    int
}
