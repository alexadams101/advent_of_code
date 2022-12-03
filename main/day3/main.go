package main
import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//Part 1
	data := loadData("main/day3/input.txt")
	total := 0
	for _, line := range data {
		duplicates := make(map[rune]int, 0)
		compartment1 := line[0:len(line)/2]
		compartment2 := line[len(line)/2:len(line)]
		for _, char := range compartment1 {
			if strings.Contains(compartment2, string(char)) {
				duplicates[char] += 1
			}
		}
		for char,_ := range duplicates {
			total += getPriority(char)
		}
	}
	fmt.Println(fmt.Sprint("Part 1: ", total))

	//Part 2
	total = 0
	for n := 0; n < len(data); n+=3 {
        group := data[n: n+3]
		badge := 'a'
		for _, char := range group[0] {
			if strings.Contains(group[1], string(char)) && strings.Contains(group[2], string(char)){
				badge = char
			}
		}
		total += getPriority(badge)
    }
	fmt.Println(fmt.Sprint("Part 2: ", total))
}

func loadData(path string) []string {
	file, _ := os.ReadFile(path)
	content := strings.Split(string(file), "\n")
	return content
}

func getPriority(char rune) int {
	if unicode.IsLower(char) {
		return int(char) - 96
	}
	return int(char) - 38
}