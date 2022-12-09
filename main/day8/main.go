package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

type Tree struct {
	Viewable bool
	Height int
}

func main() {
	//Part 1
	trees := loadData("main/day8/input.txt")
	treeCount := 0
	for _,row := range trees {
		for _, tree := range row {
			if tree.Viewable {
				treeCount++
			}
		}
	}
	fmt.Println("Part 1:", treeCount)
	//Part 2
	scores := make([]int, 0)
	for rowIdx, row := range trees {
		for colIdx, tree := range row {
			upScore := getUpScore(trees, tree.Height, colIdx, rowIdx-1, 0)
			downScore := getDownScore(trees, tree.Height, colIdx, rowIdx+1, 0)
			leftScore := getLeftScore(trees, tree.Height, colIdx-1, rowIdx, 0)
			rightScore := getRightScore(trees, tree.Height, colIdx+1, rowIdx, 0)
			score := upScore * downScore * leftScore * rightScore
			scores = append(scores, score)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	fmt.Println("Part 2:", scores[0])
}

func getUpScore(graph [][]Tree, height int, colIdx int, rowIdx int, score int) int {
	if rowIdx < 0 {
		return score
	}
	if graph[rowIdx][colIdx].Height < height{
		score++
		return getUpScore(graph, height, colIdx, rowIdx-1, score)
	}
	score++
	return score
}

func getDownScore(graph [][]Tree, height int, colIdx int, rowIdx int, score int) int {
	if rowIdx > len(graph)-1 {
		return score
	}
	if graph[rowIdx][colIdx].Height < height{
		score++
		return getDownScore(graph, height, colIdx, rowIdx+1, score)
	}
	score++
	return score
}

func getLeftScore(graph [][]Tree, height int, colIdx int, rowIdx int, score int) int {
	if colIdx < 0 {
		return score
	}
	if graph[rowIdx][colIdx].Height < height{
		score++
		return getLeftScore(graph, height, colIdx-1, rowIdx, score)
	}
	score++
	return score
}

func getRightScore(graph [][]Tree, height int, colIdx int, rowIdx int, score int) int {
	if colIdx > len(graph[0])-1 {
		return score
	}
	if graph[rowIdx][colIdx].Height < height{
		score++
		return getRightScore(graph, height, colIdx+1, rowIdx, score)
	}
	score++
	return score
}

func setInteriorViewables(trees [][]Tree) [][]Tree {
	trees[0] = setRowAsViewable(trees[0])
	trees[len(trees)-1] = setRowAsViewable(trees[len(trees)-1])
	for idx, _ := range trees[1:len(trees)-1] {
		trees[idx+1] = setColumnAsViewable(trees[idx+1])
	}
	for idx, _ := range trees[1:len(trees)-1] {
		trees[idx+1] = setViewablesFromRight(trees[idx+1], len(trees[idx+1])-2, trees[idx+1][len(trees[idx+1])-1].Height)
		trees[idx+1] = setViewablesFromLeft(trees[idx+1], 1, trees[idx+1][0].Height)
	}
	for idx, _ := range trees[0][1:len(trees)-1] {
		trees = setViewablesFromTop(trees, 1, idx+1, trees[0][idx+1].Height)
		trees = setViewablesFromBottom(trees, len(trees)-2, idx+1, trees[len(trees)-1][idx+1].Height)
	}
	return trees
}

func loadData(path string) [][]Tree {
	file, _ := os.ReadFile(path)
	rows := strings.Split(string(file), "\n")
	trees := make([][]Tree, 0)
	for _, row := range rows {
		treeRow := generateTreeRow(row)
		trees = append(trees, treeRow)
	}
	return setInteriorViewables(trees)
}

func generateTreeRow(row string) []Tree {
	trees := make([]Tree, 0)
	for _, char := range row {
		height,_ := strconv.Atoi(string(char))
		trees = append(trees, Tree{false, height})
	}
	return trees
}

func setRowAsViewable(row []Tree) []Tree{
	for idx, _ := range row {
		row[idx].Viewable = true
	}
	return row
}

func setColumnAsViewable(row []Tree) []Tree {
	row[0].Viewable = true
	row[len(row)-1].Viewable = true
	return row
}

func setViewablesFromRight(row []Tree, idx int, maxHeight int) []Tree {
	if row[idx].Height > maxHeight {
		row[idx].Viewable = true
		maxHeight = row[idx].Height
	} 
	if idx == 1 {
		return row
	}
	return setViewablesFromRight(row, idx-1, maxHeight)
}

func setViewablesFromLeft(row []Tree, idx int, maxHeight int) []Tree {
	if row[idx].Height > maxHeight {
		row[idx].Viewable = true
		maxHeight = row[idx].Height
	} 
	if idx == len(row)-2 {
		return row
	}
	return setViewablesFromLeft(row, idx+1, maxHeight)
}

func setViewablesFromTop(graph [][]Tree, idx int, column int, maxHeight int) [][]Tree {
	if graph[idx][column].Height > maxHeight {
		graph[idx][column].Viewable = true
		maxHeight = graph[idx][column].Height
	}
	if idx == len(graph)-2 {
		return graph
	}
	return setViewablesFromTop(graph, idx+1, column, maxHeight)
}

func setViewablesFromBottom(graph [][]Tree, idx int, column int, maxHeight int) [][]Tree {
	if graph[idx][column].Height > maxHeight {
		graph[idx][column].Viewable = true
		maxHeight = graph[idx][column].Height
	}
	if idx == 1 {
		return graph
	}
	return setViewablesFromBottom(graph, idx-1, column, maxHeight)
}