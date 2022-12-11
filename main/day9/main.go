package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Position struct {
	X int
	Y int
}

type Instruct struct {
	Direction string
	Quantity int
}

var tailPositions map[string]int = make(map[string]int)

func main() {
	//Part 1
	instructs := loadData("main/day9/input.txt")
	var head Position
	var tail Position
	for _, inst := range instructs { 
		for n:=0; n<inst.Quantity; n++ {
			head = simulateHead(head, inst.Direction)
			tail = simulateTail(head, tail)
			coord := fmt.Sprint(tail.X) + "," + fmt.Sprint(tail.Y)
			tailPositions[coord]++
		}
	}

	fmt.Println("Part 1:", len(tailPositions))

	//Part 2
	train := make([]Position,10)
	tailPositions = make(map[string]int)

	for _, inst := range instructs {
		for n:=0; n<inst.Quantity; n++ {
			train[0] = simulateHead(train[0], inst.Direction)
			for idx, _ := range train[1:] {
				train[idx+1] = simulateTail(train[idx], train[idx+1])
			}
			coord := fmt.Sprint(train[len(train)-1].X) + "," + fmt.Sprint(train[len(train)-1].Y)
			tailPositions[coord]++
		}
	}
	fmt.Println("Part 2:", len(tailPositions))

}

func simulateHead(head Position, direction string) Position{
	switch direction {
	case "U":
		head.Y++
	case "D":
		head.Y--
	case "R":
		head.X++
	case "L":
		head.X--
	}
	return head
}

func simulateTail(head Position, tail Position) Position {
	xDiff := head.X-tail.X
	yDiff := head.Y-tail.Y
	tailXMove := calculateMove(xDiff)
	tailYMove := calculateMove(yDiff)

	if tailXMove !=0 && tailYMove ==0 && yDiff !=0 {
		tailYMove += yDiff
	} else if tailYMove != 0 && tailXMove ==0 && xDiff !=0 {
		tailXMove += xDiff
	}

	tail.X += tailXMove
	tail.Y += tailYMove

	return tail
}

func calculateMove(diff int) int {
	if diff > 0 {
		return diff - 1
	} 
	if diff < 0 {
		return diff + 1
	}
	return 0
}

func loadData(path string) []Instruct {
	file, _ := os.ReadFile(path)
	lines := strings.Split(string(file), "\n")
	instructs := make([]Instruct, 0)
	for _, line := range lines {
		i := strings.Split(line, " ")
		quant,_ := strconv.Atoi(i[1])
		instructs = append(instructs, Instruct{i[0],quant})
	}
	return instructs
}