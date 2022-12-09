package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

type File struct {
	Type string
	Parent *File
	Children map[string]*File
	Size int
}
func (file *File) set_parent_size() {
	if file.Parent != nil {
		file.Parent.Size += file.Size
	}
}
func create_dir(parent *File) *File {
	return &File{"DIRECTORY", parent, make(map[string]*File), 0}
}
func create_file(parent *File, size int) *File {
	return &File{"FILE", parent, make(map[string]*File), size}
}

func main() {
	//Part 1
	files := loadData("main/day7/test_input.txt")
	fileSizes := []int{}
	for i:=len(files)-1; i>=0; i-- {
		for _, file := range files[i]{
			file.set_parent_size()
		}
	}
	for _, row := range files {
		for _, file :=range row {
			if file.Size<=100000 && file.Type=="DIRECTORY" {
				fileSizes = append(fileSizes,file.Size)
			}
		}
	}
	total := 0
	for _, size := range fileSizes {
		total+=size
	}
	fmt.Println("Part 1:", total)

	deletableFiles := []int{}
	deletableBar := 30000000-(70000000-files[0][0].Size)
	for _, row := range files {
		for _, file :=range row {
			if file.Size>=deletableBar && file.Type=="DIRECTORY" {
				deletableFiles = append(deletableFiles,file.Size)
			}
		}
	}
	sort.Sort(sort.IntSlice(deletableFiles))
	fmt.Println("Part 2:", deletableFiles[0])
}

func loadData(path string) [][]*File {
	file, _ := os.ReadFile(path)
	files := make([][]*File,0)
	content := strings.Split(string(file), "\n")
	root := create_dir(nil)
	currentDir := root
	dirLevel := 0
	files = append(files, []*File{root})
	dirLevel++
	for _, line := range content[1:] {
		if line[0:] == "$ cd .." {
			dirLevel --
			currentDir = currentDir.Parent
		} else if line[0:4] == "$ cd" {
			dirLevel ++
			dirName := line[5:]
			currentDir = currentDir.Children[dirName]
		} else if line[0:4] == "$ ls" {
		} else if line[0:3] == "dir" {
			if dirLevel>len(files)-1 {
				files = append(files, []*File{})
			}
			dirName := line[4:]
			dir := create_dir(currentDir)
			currentDir.Children[dirName] = dir
			files[dirLevel] = append(files[dirLevel], dir)
		} else {
			if dirLevel>len(files)-1 {
				files = append(files, []*File{})
			}
			file := strings.Split(line, " ")
			name := file[1]
			size,_ := strconv.Atoi(file[0])
			f := create_file(currentDir, size)
			currentDir.Children[name] = f
			files[dirLevel] = append(files[dirLevel], f)
		}
	}
	return files
}