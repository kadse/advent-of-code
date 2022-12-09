package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

const (
	//objectives
	rockV1        = "A"
	paperV1       = "B"
	scissorsV1    = "C"
	rockV2        = "X"
	paperV2       = "Y"
	scissorsV2    = "Z"
	rock          = 0
	paper         = 1
	scissors      = 2
	scoreRock     = 1
	scorePaper    = 2
	scoreScissors = 3

	//scoring
	lose      = 0
	draw      = 1
	win       = 2
	scoreLose = 0
	scoreDraw = 3
	scoreWin  = 6
)

func main() {
	part1()
	fmt.Println()
	part2()
}

func part1() {
	fmt.Println("Example Part 1:", execPart1(true)) //15
	fmt.Println("Part 1:", execPart1(false))        //15572
}

func part2() {
	fmt.Println("Example Part 2:", execPart2(true)) //12
	fmt.Println("Part 2:", execPart2(false))
}

func execPart1(example bool) int {
	score := 0
	for _, round := range prepareData(example) {
		choose := strings.Fields(round)
		score += getScore(choose[1], choose[0])
	}

	return score
}

func execPart2(example bool) int {
	score := 0
	for _, round := range prepareData(example) {
		choose := strings.Fields(round)
		score += getScore(choose[1], choose[0])
	}

	return score
}

func prepareData(example bool) []string {
	return strings.Split(utils.GetData(example), "\n")
}

func isRock(objective string) bool {
	return objective == rockV1 || objective == rockV2
}

func isPaper(objective string) bool {
	return objective == paperV1 || objective == paperV2
}

func isScissors(objective string) bool {
	return objective == scissorsV1 || objective == scissorsV2
}

func getScore(myChoose string, opponentChoose string) int {
	resultMap := [3][3]int{
		{draw, lose, win},
		{win, draw, lose},
		{lose, win, draw},
	}

	score := 0
	switch resultMap[getObjetiveIndex(myChoose)][getObjetiveIndex(opponentChoose)] {
	case lose:
		score += scoreLose
	case draw:
		score += scoreDraw
	case win:
		score += scoreWin
	}

	return score + getScoreObjective(myChoose)
}

func getObjetiveIndex(objective string) int {
	switch true {
	case isRock(objective):
		return rock
	case isPaper(objective):
		return paper
	case isScissors(objective):
		return scissors
	}

	return -1
}

func getScoreObjective(objective string) int {
	switch true {
	case isRock(objective):
		return scoreRock
	case isPaper(objective):
		return scorePaper
	case isScissors(objective):
		return scoreScissors
	}

	return -1
}
