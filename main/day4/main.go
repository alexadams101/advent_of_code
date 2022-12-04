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

type fits func([]int, []int) bool

func main() {
	//Part 1
	data := loadData("main/day4/input.txt")
	noOfPairs := 0
	for _, pair := range data {
		if containsSlice(pair.section1, pair.section2, sliceFitsComplete) {
			noOfPairs++
		}
	}
	fmt.Println("Part 1:", noOfPairs)

	//Part 2
	noOfPairs = 0
	for _, pair := range data {
		if containsSlice(pair.section1, pair.section2, sliceFitsPartial) {
			noOfPairs++
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

func containsSlice(slice1 []int, slice2 []int, fits fits) bool {
	if len(slice1) > len(slice2) && fits(slice1, slice2) {
		return true
	} 
	if len(slice2) >= len(slice1) && fits(slice2, slice1) {
		return true
	}
	return false
}

func sliceFitsComplete(slice []int, sliceToFit []int) bool {
	for _, v := range sliceToFit {
		if !contains(slice,v) {
			return false
		}
	}
	return true
}

func sliceFitsPartial(slice []int, sliceToFit []int) bool {
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