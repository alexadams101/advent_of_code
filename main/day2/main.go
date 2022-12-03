package main
import (
	"fmt"
	"os"
	"strings"
)

type move int64
type result int64
type roundPart1 struct {
	opponentMove move
	yourMove move
}
type roundPart2 struct {
	opponentMove move
	result result
}

const (
	Rock move = 1
	Paper move = 2
	Scissors move = 3
)
const (
	Loss result = 0
	Draw result = 3
	Win result = 6
)

func main() {
	//Part 1
	data := loadData("main/day2/input.txt")

	roundsPart1 := make([]roundPart1,0)

	for _, line := range data {
		roundsPart1 = append(roundsPart1, roundPart1{opponentMove: getMove(line[0]), yourMove: getMove(line[1])})
	}

	totalPart1 := 0
	for _, round := range roundsPart1 {
		result := generateResult(round.opponentMove, round.yourMove)
		totalPart1 += int(result) + int(round.yourMove)
	}

	fmt.Println(fmt.Sprint("Part 1: ", totalPart1))

	//Part 2
	roundsPart2 := make([]roundPart2,0)
	
	for _, line := range data {
		roundsPart2 = append(roundsPart2, roundPart2{opponentMove: getMove(line[0]), result: getResult(line[1])})
	}

	totalPart2 := 0
	for _, round := range roundsPart2 {
		move := findMove(round.opponentMove, round.result)
		totalPart2 += int(round.result) + int(move)
	}

	fmt.Println(fmt.Sprint("Part 2: ", totalPart2))
}

func loadData(path string) [][]string {
	file, _ := os.ReadFile(path)
	content := strings.Split(string(file), "\n")
	
	list := make([][]string, 0)
	for _, line := range content {
		vals := strings.Split(line, " ")
		list = append(list, vals)
	}
	return list
}

func getMove(val string) move {
	switch val {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	default:
		return Scissors
	}
}

func getResult(val string) result {
	switch val {
	case "X":
		return Loss
	case "Y":
		return Draw
	default:
		return Win
	}
}

func generateResult(opponentMove move, yourMove move) result {
	switch opponentMove {
	case Rock:
		switch yourMove {
		case Paper:
			return Win
		case Scissors:
			return Loss
		}
	case Paper:
		switch yourMove {
		case Rock:
			return Loss
		case Scissors:
			return Win
		}
	case Scissors:
		switch yourMove {
		case Rock:
			return Win
		case Paper:
			return Loss
		}
	}
	return Draw
}

func findMove(opponentMove move, result result) move {
	switch opponentMove {
	case Rock:
		switch result {
		case Loss:
			return Scissors
		case Win:
			return Paper
		}
	case Paper:
		switch result {
		case Loss:
			return Rock
		case Win:
			return Scissors
		}	
	case Scissors:
		switch result {
		case Loss:
			return Paper
		case Win:
			return Rock
		}
	}
	return opponentMove
}