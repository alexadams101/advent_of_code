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
	data := loadData("main/day1/input.txt")

	totals := make([]int, 0)
	for _, elf := range data {
		totals = append(totals, sum(elf))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	fmt.Println(totals[0])

	//Part 2
	top3 := sum([]int {totals[0], totals[1], totals[2]})
	fmt.Println(top3)
}

func loadData(path string) [][]int {
	file, _ := os.ReadFile(path)
	content := strings.Split(string(file), "\n")

	list := make([][]int, 0)
	elf := make([]int, 0)
    for idx, val := range content {
		if val != "" {
			calories, _ := strconv.Atoi(val)
			elf = append(elf, calories)	
		} 
		if idx==len(content)-1 {
			list = append(list, elf)
		}
		if val == "" {
			list = append(list, elf)
			elf = make([]int, 0)
		}
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