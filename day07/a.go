package day07

import (
	"strings"
)

var fs = FileSystem{}

var dirSizes int

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "07a"
}

func (p PuzzleA) Run() any {
	inputSlice := strings.Split(input, "\n")
	fs.InitFileSystem(&inputSlice)
	fs.ReadSizeRecursive()
	return dirSizes
}

func (fs *FileSystem) ReadSizeRecursive() {
	if fs.currentDir == nil {
		return
	}

	has, dir := (*fs).currentDir.HasDirWithNoSizes()
	if has {
		fs.currentDir = dir
		fs.ReadSizeRecursive()
	} else {
		fs.currentDir.CountDirSize()
		fs.currentDir = fs.currentDir.parentDir
		fs.ReadSizeRecursive()
	}
}

// CountDirSize counts the size of all files in the current directory,
// assigns this value to the directory and if it's not more than 100000,
// it adds that value to the global dir-size register.
func (d *Directory) CountDirSize() {

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
	if dirSize <= 100000 {
		dirSizes += dirSize
	}
}
