package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type (
	Filesystem struct {
		root *Directory
	}

	Directory struct {
		name           string
		prevDirectory  *Directory
		subDirectories []*Directory
		files          []*File
	}

	File struct {
		name string
		size int
	}
)

var (
	TotalDiskSpace = 70000000
	SpaceNeeded    = 30000000
	MinimumSize    = 100000
)

func day7() {
	file, err := os.Open("../inputs/day7.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open input file for day 7: %v", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("Failed to close input file for day 7: %v", err))
		}
	}(file)

	var fs = Filesystem{root: &Directory{
		name: "/",
	}}

	var currentDir *Directory

	isLs := false

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		cmd := strings.Split(s, " ")
		if cmd[0] == "$" {
			command := cmd[1]
			isLs = false

			switch command {
			case "cd":
				currentDir = fs.ChangeDir(cmd[2], currentDir)
				break

			case "ls":
				isLs = true
				break

			default:
				fmt.Printf("Don't know what to do with cmd \"%s\"", command)
			}
		} else {
			if isLs {
				if cmd[0] == "dir" {
					currentDir.subDirectories = append(currentDir.subDirectories, &Directory{
						name:          cmd[1],
						prevDirectory: currentDir,
					})
				} else {
					currentDir.files = append(currentDir.files, &File{
						name: cmd[1],
						size: parseInt(cmd[0]),
					})
				}
			}
		}
	}

	resPuzzle1 := findResultDay7Puzzle1(&fs)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 1 at Day 7: %d", resPuzzle1))

	resPuzzle2 := findResultDay7Puzzle2(&fs)
	fmt.Println(fmt.Sprintf("Answer for Puzzle 2 at Day 7: %d", resPuzzle2))
}

func findResultDay7Puzzle1(fs *Filesystem) int {
	//fs.Dump()

	dirSizes := map[string]int{
		fs.root.name: fs.root.CalculateSizeRecursively(),
	}

	for _, directory := range fs.root.subDirectories {
		addDirSizes(dirSizes, directory, "/")
	}

	sortMapByValues(dirSizes)

	dirSizesAtMost100000 := make([]int, len(dirSizes))
	for _, dirSize := range dirSizes {
		//fmt.Printf("%s: %d\n", dirName, dirSize)

		if dirSize <= MinimumSize {
			//fmt.Printf("%s: %d\n", dirName, dirSize)
			dirSizesAtMost100000 = append(dirSizesAtMost100000, dirSize)
		}
	}

	return sum(dirSizesAtMost100000...)
}

func findResultDay7Puzzle2(fs *Filesystem) int {
	dirSizes := map[string]int{
		//fs.root.name: fs.root.CalculateSizeRecursively(),
	}

	rootDirSize := fs.root.CalculateSizeRecursively()
	for _, directory := range fs.root.subDirectories {
		addDirSizesIfWorthDeleting(dirSizes, directory, rootDirSize, "/")
	}

	//sortMapByValues(dirSizes)

	smallestDirSize := rootDirSize
	for _, dirSize := range dirSizes {
		//fmt.Printf("%s: %d\n", dirName, dirSize)
		smallestDirSize = min(smallestDirSize, dirSize)
	}

	return smallestDirSize
}

func addDirSizes(dirSizes map[string]int, dir *Directory, parent string) {
	dirSizes[parent+dir.name] = dir.CalculateSizeRecursively()

	if dir.subDirectories != nil {
		for _, subDir := range dir.subDirectories {
			addDirSizes(dirSizes, subDir, parent+"/"+dir.name)
		}
	}
}

func addDirSizesIfWorthDeleting(dirSizes map[string]int, dir *Directory, usedSpace int, parent string) {
	dirSize := dir.CalculateSizeRecursively()
	if TotalDiskSpace-usedSpace+dirSize > SpaceNeeded {
		dirSizes[parent+dir.name] = dirSize
	}

	if dir.subDirectories != nil {
		for _, subDir := range dir.subDirectories {
			addDirSizesIfWorthDeleting(dirSizes, subDir, usedSpace, parent+"/"+dir.name)
		}
	}
}

func (f *Filesystem) ChangeDir(path string, currentDir *Directory) *Directory {
	if path == "/" {
		return f.root
	}

	if path == ".." {
		return currentDir.prevDirectory
	}

	for _, subDirectory := range currentDir.subDirectories {
		if subDirectory.name == path {
			return subDirectory
		}
	}

	return nil
}

func (d *Directory) CalculateSizeRecursively() int {
	totalSize := 0

	for _, file := range d.files {
		totalSize += file.size
	}

	if d.subDirectories != nil {
		for _, subDirectory := range d.subDirectories {
			totalSize += subDirectory.CalculateSizeRecursively()
		}
	}

	return totalSize
}

func (f *Filesystem) Dump() {
	lines := []string{
		fmt.Sprintf("- %s (dir)", f.root.name),
	}

	for _, subDirectory := range f.root.subDirectories {
		addLineDirectory(&lines, subDirectory, 1)
	}

	for _, file := range f.root.files {
		addLineFile(&lines, file, 1)
	}

	for _, line := range lines {
		println(line)
	}
}

func addLineDirectory(lines *[]string, d *Directory, depth int) {
	*lines = append(*lines, fmt.Sprintf("%s- %s (dir)", strings.Repeat(" ", depth*2), d.name))

	if d.subDirectories != nil {
		for _, subDirectory := range d.subDirectories {
			addLineDirectory(lines, subDirectory, depth+1)
		}
	}

	if d.files != nil {
		for _, file := range d.files {
			addLineFile(lines, file, depth+1)
		}
	}
}

func addLineFile(lines *[]string, f *File, depth int) {
	*lines = append(*lines, fmt.Sprintf("%s- %s (file, size=%d)", strings.Repeat(" ", depth*2), f.name, f.size))
}
