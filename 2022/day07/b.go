package day07

import (
	"strings"
)

var dirMap = make(map[string]int)

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "07b"
}

func (p PuzzleB) Run() any {
	var fs = FileSystem{}
	inputSlice := strings.Split(input, "\n")
	fs.InitFileSystem(&inputSlice)
	fs.ReadSizeRecursive2()

	return GetLowestValue(dirMap)
}

func GetLowestValue(sl map[string]int) int {
	lowest := 70000000
	for _, val := range sl {
		if val < lowest {
			lowest = val
		}
	}
	return lowest
}

func (fs *FileSystem) ReadSizeRecursive2() {
	if fs.currentDir == nil {
		return
	}

	has, dir := (*fs).currentDir.HasDirWithNoSizes()
	if has {
		fs.currentDir = dir
		fs.ReadSizeRecursive2()
	} else {
		fs.currentDir.CountDirSize2()
		fs.currentDir = fs.currentDir.parentDir
		fs.ReadSizeRecursive2()
	}
}

func (d *Directory) CountDirSize2() {
	dirSize := 0
	for _, val := range (*d).files {
		dirSize += val.size
	}

	if (*d).childrenDir != nil {
		for _, val := range (*d).childrenDir {
			dirSize += val.size
		}
	}

	(*d).size = dirSize

	if dirSize >= 3636666 {
		dirMap[(*d).name] = dirSize
	}
}
