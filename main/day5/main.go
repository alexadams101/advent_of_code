package main
import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

type Instruction struct {
	noOfCrates int
	sourceStack int
	targetStack int
}

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) PushMulti(strings []string) {
	for _,str := range strings {
		*s = append(*s, str)
	}
}

func (s *Stack) Pop() string {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

func (s *Stack) PopMulti(no int) []string {
	index := len(*s) - no
	element := (*s)[index:]
	*s = (*s)[:index]
	return element
}

func main() {
	//Part 1
	stacks, instructions := loadData("main/day5/test_input.txt")
	for _,inst := range instructions {
		for n:=0; n<inst.noOfCrates; n++{
			poppedStr := stacks[inst.sourceStack-1].Pop()
			stacks[inst.targetStack-1].Push(poppedStr)
		}
	}
	output := ""
	for _, stack := range stacks {
		output += stack[len(stack)-1]
	}
	fmt.Println("Part 1:", output)

	//Part 2
	stacks, instructions = loadData("main/day5/input.txt")
	for _,inst := range instructions {
		poppedStrings := stacks[inst.sourceStack-1].PopMulti(inst.noOfCrates)
		stacks[inst.targetStack-1].PushMulti(poppedStrings)
	}
	outputPt2 := ""
	for _, stack := range stacks {
		outputPt2 += stack[len(stack)-1]
	}
	fmt.Println("Part 2:", outputPt2)
}

func loadData(path string) ([]Stack,[]Instruction) {
	file, _ := os.ReadFile(path)
	sections := strings.Split(string(file), "\n\n")
	stks := strings.Split(sections[0], "\n")
	stacks := make([]Stack, 0)
	for r:=1; r<len(stks[len(stks)-1]); r+=4 {
		var stack Stack
		for c:=len(stks)-2; c>=0; c-- {
			if string(stks[c][r]) != " " {
				stack.Push(string(stks[c][r]))
			}
		}
		stacks = append(stacks, stack)
	}

	instructs := strings.Split(sections[1], "\n")
	instructions := make([]Instruction, 0)
	for _, instruct := range instructs {
		digits := strings.Split(instruct, " ")
		number, _ := strconv.Atoi(digits[1])
		source, _ := strconv.Atoi(digits[3])
		target, _ := strconv.Atoi(digits[5])
		instructions = append(instructions, Instruction{number, source, target})
	}
	return stacks, instructions
}