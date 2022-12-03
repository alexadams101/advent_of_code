package main
import (
    "fmt"
    "os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	//Part 1
	elves := loadData("main/day1/input.txt")

	totals := make([]int, 0)
	for _, elf := range elves {
		totals = append(totals, sum(elf))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	fmt.Println(fmt.Sprint("Part 1: ", totals[0]))

	//Part 2
	top3 := totals[0:3]
	fmt.Println(fmt.Sprint("Part 2: ", sum(top3)))
}

func loadData(path string) [][]int {
	file, _ := os.ReadFile(path)
	content := strings.Split(string(file), "\n\n")

	list := make([][]int, 0)
	for _, val := range content {
		elf := make([]int, 0)
		for _,calorie := range strings.Split(val, "\n") {
			calories, _ := strconv.Atoi(calorie)
			elf = append(elf, calories)	
		}
		list = append(list, elf)
	}
	return list
}

func sum(list []int) int {
	result := 0  
	for _, v := range list {  
		result += v  
	}  
	return result
}