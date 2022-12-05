package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	fmt.Println()
	part2()
}

func part1() {
	fmt.Println("Example Part 1:", execPart1(true))
	fmt.Println("Part 1:", execPart1(false))
}

func part2() {
	fmt.Println("Example Part 2:", execPart2(true))
	fmt.Println("Part 2:", execPart2(false))
}

func execPart1(example bool) int {
	maxCalories := 0
	for _, elveCalories := range getData(example) {
		sumCalories := getCaloriesByElve(elveCalories)

		if maxCalories < sumCalories {
			maxCalories = sumCalories
		}
	}

	return maxCalories
}

func execPart2(example bool) int {
	var calories []int
	for _, elveCalories := range getData(example) {
		calories = append(calories, getCaloriesByElve(elveCalories))
	}

	sort.Ints(calories)
	lengthSlice := len(calories)

	sumCalories := 0
	for i := lengthSlice - 1; i >= (lengthSlice - 3); i-- {
		sumCalories += calories[i]
	}

	return sumCalories
}

func getCaloriesByElve(elveCalories string) int {
	sumCalories := 0
	for _, calories := range strings.Fields(elveCalories) {
		intCalories, _ := strconv.Atoi(calories)
		sumCalories += intCalories
	}

	return sumCalories
}

func getData(example bool) []string {
	return strings.Split(utils.GetData(example), "\n\n")
}
