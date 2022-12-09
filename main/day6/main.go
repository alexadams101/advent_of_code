package main
import (
	"fmt"
	"os"
)

func main() {
	//Part 1
	text := loadData("main/day6/input.txt")
	var output int
	for n:=0; n<len(text)-4; n++{
		marker := text[n:n+4]
		if !thereIsDuplicate(marker) {
			output = n+4
			break
		}
	}
	fmt.Println("Part 1:", output)

	//Part 2
	var outputPt2 int
	for n:=0; n<len(text)-14; n++{
		marker := text[n:n+14]
		if !thereIsDuplicate(marker) {
			outputPt2 = n+14
			break
		}
	}
	fmt.Println("Part 2:", outputPt2)
}

func loadData(path string) string {
	file, _ := os.ReadFile(path)
	return string(file)
}

func thereIsDuplicate(marker string) bool {
	m := make(map[rune]int)
	for _,letter := range marker {
		m[letter] ++
	}
	if len(m)<len(marker) {
		return true
	}
	return false
}

