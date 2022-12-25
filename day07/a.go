package day07

import (
	"strconv"
	"strings"
)

var fs = FileSystem{}

var dirSizes int

type PuzzleA struct{}

func (p PuzzleA) String() string {
	return "07a"
}

type File struct {
	name string
	size int
}

type Directory struct {
	name        string
	size        int
	files       []*File
	childrenDir []*Directory
	parentDir   *Directory
}

type FileSystem struct {
	currentDir *Directory
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

func (d *Directory) HasDirWithNoSizes() (bool, *Directory) {

	if (*d).childrenDir == nil {
		return false, nil
	}

	for _, dir := range (*d).childrenDir {
		if (*dir).size == 0 {
			return true, dir
		}
	}

	return false, nil
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

func (fs *FileSystem) InitFileSystem(inputSlice *[]string) {
	for _, val := range *inputSlice {
		line := strings.Split(val, " ")
		switch line[1] {
		case "cd":
			fs.ChangeHandler(&line)
		case "ls":
			continue
		default:
			fs.MakeHandler(&line)
		}
	}
}

func (fs *FileSystem) MakeHandler(line *[]string) {
	switch (*line)[0] {
	case "dir":
		(*fs.currentDir).MkDir(line)
	default:
		(*fs.currentDir).MkFile(line)
	}
}

func (fs *FileSystem) ChangeHandler(line *[]string) {
	arg := (*line)[2]
	switch arg {
	case "/":
		root := Directory{
			name: "root",
			size: 0,
		}
		fs.currentDir = &root
	case "..":
		fs.currentDir = fs.currentDir.parentDir
	default:
		dirs := fs.currentDir.childrenDir
		for i := 0; i < len(dirs); i++ {
			if dirs[i].name == arg {
				fs.currentDir = dirs[i]
			}
		}
	}
}

func (d *Directory) MkFile(line *[]string) {
	size, _ := strconv.Atoi((*line)[0])
	file := File{
		name: (*line)[1],
		size: size,
	}
	d.files = append(d.files, &file)
}

func (d *Directory) MkDir(line *[]string) {
	dir := Directory{
		name:      (*line)[1],
		parentDir: d,
	}
	d.childrenDir = append(d.childrenDir, &dir)
}
