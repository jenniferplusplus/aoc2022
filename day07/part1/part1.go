package part1

import (
	"strconv"
	"strings"
)

type File struct {
	Size int32
	Name string
}

type Dir struct {
	Name   string
	Root   *Dir
	Parent *Dir
	Links  map[string]*Dir
	Files  map[string]*File
}

func (cwd *Dir) AddDir(name string) *Dir {
	child := NewDir(name, cwd.Root)
	child.Parent = cwd
	cwd.Links[name] = child

	return child
}

func (cwd *Dir) AddFile(name string, size int32) *File {
	file := &File{Name: name, Size: size}
	cwd.Files[name] = file

	return file
}

func (cwd *Dir) Size(out chan int64) int64 {
	if out != nil {
		defer close(out)
	}
	return cwd._size(out)
}

func NewRoot() *Dir {
	root := NewDir("", nil)
	root.Root = root
	return root
}

func NewDir(name string, root *Dir) *Dir {
	links := map[string]*Dir{}
	files := map[string]*File{}
	dir := &Dir{
		Name:  name,
		Root:  root,
		Links: links,
		Files: files,
	}
	return dir
}

func Parse(lines []string) *Dir {
	cwd := NewRoot()
	root := cwd.Root
	for _, line := range lines {
		cwd = parseLine(line, cwd)
	}

	return root
}

func (cwd *Dir) _size(out chan int64) int64 {
	size := int64(0)
	for _, dir := range cwd.Links {
		size += dir._size(out)
	}

	for _, file := range cwd.Files {
		size += int64(file.Size)
	}

	if out != nil {
		out <- size
	}

	return size
}

func parseLine(line string, cwd *Dir) *Dir {
	tokens := strings.Split(line, " ")
	switch tokens[0] {
	case "$":
		switch tokens[1] {
		case "ls":
		default:
			switch tokens[2] {
			case "/":
				cwd = cwd.Root
			case "..":
				cwd = cwd.Parent
			default:
				cwd = cwd.Links[tokens[2]]
			}
		}
	case "dir":
		cwd.AddDir(tokens[1])
	default:
		size, _ := strconv.ParseInt(tokens[0], 10, 32)
		cwd.AddFile(tokens[1], int32(size))
	}

	return cwd
}
