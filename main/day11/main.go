package main
import (
	"fmt"
	"os" 
	"strings"
	"strconv"
	"sort"
)

type Operation func(int) int

type Test func(int) bool

type Div func(int) int

type Monkey struct {
	Items []int
	Operation Operation
	Test int
	TrueMonkey int
	FalseMonkey int
}

var monkeyInspections []int

var trials []Div

func main() {
	//Part 1
	monkeys := loadData("main/day11/input.txt")
	monkeyInspections = make([]int, len(monkeys))

	for round:=0;round<20;round++{
		for idx:=0;idx<len(monkeys);idx++ {
			monkeyInspections[idx] += len(monkeys[idx].Items)
			monkeys = simulate(idx, monkeys, func (i int) int {return i/3})
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(monkeyInspections)))
	product := monkeyInspections[0] * monkeyInspections[1]
	fmt.Println("Part 1:", product)

	//Part 2
	monkeys = loadData("main/day11/input.txt")
	monkeyInspections = make([]int, len(monkeys))

	mod := monkeys[0].Test

	for _, m := range monkeys[1:] {
		mod *= m.Test
	}

	for round:=0;round<10000;round++{
		for idx:=0;idx<len(monkeys);idx++ {
			monkeyInspections[idx] += len(monkeys[idx].Items)
			monkeys = simulate(idx, monkeys, func (i int) int {return i % mod})
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(monkeyInspections)))
	product = monkeyInspections[0] * monkeyInspections[1]
	fmt.Println("Part 2:", product)
}

func loadData(path string) []Monkey {
	file,_ := os.ReadFile(path)
	monkeyDetails := strings.Split(string(file), "\n\n")
	monkeys := make([]Monkey,0)
	for _,monkey := range monkeyDetails {
		monkeyInfo := strings.Split(string(monkey), "\n")
		newMonkey := Monkey {
			generateItems(monkeyInfo[1]),
			generateOperation(monkeyInfo[2]),
			generateTest(monkeyInfo[3]),
			generateMonkey(monkeyInfo[4]),
			generateMonkey(monkeyInfo[5]),
		}
		monkeys = append(monkeys, newMonkey)
	}
	return monkeys
}

func generateItems(line string) []int {
	startingItems := strings.Split (line[18:], ", ")
	integerItems := make([]int, 0)
	for _, item := range startingItems {
		val,_ := strconv.Atoi(item)
		integerItems = append(integerItems, val)
	}
	return integerItems
}

func generateOperation(line string) Operation {
	operationParams := strings.Split (line[23:], " ")
	if operationParams[1] == "old" {
		switch operationParams[0] {
		case "*":
			return func(old int) int { return old * old }
		case "+":
			return func(old int) int { return old + old }
		case "-":
			return func(old int) int { return old - old }
		case "/":
			return func(old int) int { return old / old }
		}
	}
	val,_ := strconv.Atoi(operationParams[1])
	switch operationParams[0] {
	case "*":
		return func(old int) int { return old * val }
	case "+":
		return func(old int) int { return old + val }
	case "-":
		return func(old int) int { return old - val }
	default:
		return func(old int) int { return old / val }
	}
}

func generateTest(line string) int {
	val,_ := strconv.Atoi(line[21:])
	return val
}

func generateMonkey(line string) int {
	words := strings.Split (line[4:], " ")
	val,_ := strconv.Atoi(words[5])
	return val
}

func simulate(idx int, monkeys []Monkey, div Div) []Monkey{
	for i:=0; i<len(monkeys[idx].Items); i++ {
		output := div(monkeys[idx].Operation(monkeys[idx].Items[i]))
		if output % monkeys[idx].Test == 0 {
			monkeys[monkeys[idx].TrueMonkey].Items = append(monkeys[monkeys[idx].TrueMonkey].Items, output)
		}else {
			monkeys[monkeys[idx].FalseMonkey].Items = append(monkeys[monkeys[idx].FalseMonkey].Items, output)
		}
	}
	monkeys[idx].Items = make([]int,0)
	return monkeys
}