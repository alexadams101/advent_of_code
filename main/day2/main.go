package main
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//Part 1
	data := loadData("main/day2/input.txt")
	totalPt1 := 0
	for _, round := range data {
		winningVal := findWinner(round[0], round[1])
		value := getValue(round[1])
		switch winningVal{
		case "X", "Y", "Z":
			totalPt1 += value + 6
		case "":
			totalPt1 += value + 3
		default:
			totalPt1 += value
		}
	}
	fmt.Println(fmt.Sprint("Part 1: ", totalPt1))

	//Part 2
	totalPt2 := 0
	for _, round := range data {
		value := getValue(findValue(round[0], round[1]))
		switch round[1]{
		case "Z":
			totalPt2 += value + 6
		case "Y":
			totalPt2 += value + 3
		case "X":
			totalPt2 += value
		}
	}
	fmt.Println(fmt.Sprint("Part 2: ", totalPt2))
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

func getValue(val string) int {
	if val == "A" || val == "X" {
		return 1
	}
	if val == "B" || val == "Y" {
		return 2
	}
	return 3
}

func findWinner(val1 string, val2 string) string {
	switch val1 {
	case "A":
		switch val2 {
		case "Y":
			return "Y"
		case "Z":
			return "A"
		}
	case "B":
		switch val2 {
		case "X":
			return "B"
		case "Z":
			return "Z"
		}
	case "C":
		switch val2 {
		case "X":
			return "X"
		case "Y":
			return "C"
		}
	}
	return ""
}

func findValue(val1 string, val2 string) string {
	switch val1 {
	case "A":
		switch val2 {
		case "X":
			return "Z"
		case "Z":
			return "Y"
		default:
			return "X"
		}
	case "B":
		switch val2 {
		case "X":
			return "X"
		case "Z":
			return "Z"
		default:
			return "Y"
		}	
	case "C":
		switch val2 {
		case "X":
			return "Y"
		case "Z":
			return "X"
		default:
			return "Z"
		}
	}
	return ""
}