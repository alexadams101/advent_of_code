package main
import (
	"fmt"
	"strings"
	"os"
	"strconv"
)

func main() {
	//Part 1
	content := loadData("main/day10/input.txt")
	cycles := make([]int, 0)
	idx := 0
	currentValue := 1
	for _,line := range content {
		if line[:4] == "noop"{
			cycles = append(cycles, currentValue)
			idx++
		}
		if line[:4] == "addx"{
			v,_ := strconv.Atoi(line[5:])
			cycles = append(cycles, currentValue)
			cycles = append(cycles, currentValue)
			idx +=2
			currentValue += v
		}
	}

	cyclesToMeasure := [6]int{20,60,100,140,180,220}
	signalTotal := 0

	for _, cycleNo := range cyclesToMeasure{
		signalTotal += cycles[cycleNo-1]*cycleNo
	}
	fmt.Println("Part 1:", signalTotal)

	output := make([]string, 1)
	lineCount := 0
	charCount := 0

	for _, cycle := range cycles {
		if charCount > cycle-2 && charCount < cycle+2 {
			output[lineCount] += "#"
		} else {
			output[lineCount] += "."
		}

		charCount ++

		if charCount == 40 {
			output = append(output, "")
			lineCount ++
			charCount = 0
		}

		if lineCount == 1 {
		}
	}

	fmt.Println("Part 2:")
	for _,line := range output {
		fmt.Println(line)
	}
}

func loadData(path string) []string {
	file,_ := os.ReadFile(path)
	content := strings.Split(string(file), "\n")
	return content
}