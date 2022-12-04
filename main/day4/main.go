package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

type pair struct {
	section1 []int
	section2 []int
}

func main() {
	//Part 1
	data := loadData("main/day4/input.txt")
	noOfPairs := 0
	for _, pair := range data {
		if len(pair.section1) > len(pair.section2) && containsSlicePart1(pair.section1, pair.section2) {
			noOfPairs ++
		} 
		if len(pair.section2) >= len(pair.section1) && containsSlicePart1(pair.section2, pair.section1) {
			noOfPairs ++
		}
	}
	fmt.Println("Part 1:", noOfPairs)

	//Part 2
	noOfPairs = 0
	for _, pair := range data {
		if len(pair.section1) > len(pair.section2) && containsSlicePart2(pair.section1, pair.section2) {
			noOfPairs ++
		} 
		if len(pair.section2) >= len(pair.section1) && containsSlicePart2(pair.section2, pair.section1) {
			noOfPairs ++
		}
	}
	fmt.Println("Part 2:", noOfPairs)
}

func loadData(path string) []pair {
	file, _ := os.ReadFile(path)
	content := strings.Split(string(file), "\n")
	data := make([]pair,0)
	for _, line := range content {
		elves := strings.Split(line, ",")
		sections := pair{getSections(elves[0]), getSections(elves[1])}
		data = append(data, sections)
	}
	return data
}

func getSections(ref string) []int {
	sectionRange := strings.Split(ref, "-")
	sections := make([]int,0)
	min, _ := strconv.Atoi(sectionRange[0])
	max, _ := strconv.Atoi(sectionRange[1])
	for n:=min; n<=max; n++ {
		sections = append(sections,n)
	}
	return sections
}

func containsSlicePart1(slice []int, sliceToFit []int) bool {
	for _, v := range sliceToFit {
		if !contains(slice,v) {
			return false
		}
	}
	return true
}

func containsSlicePart2(slice []int, sliceToFit []int) bool {
	for _, v := range sliceToFit {
		if contains(slice,v) {
			return true
		}
	}
	return false
}

func contains(slice []int, str int) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}