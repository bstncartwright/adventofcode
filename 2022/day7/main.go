package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	tree := parse(input)

	// update needs 30000000
	// max size is  70000000
	rootSize := tree.Size()
	available := 70000000 - rootSize
	needed := 30000000 - available

	part1, part2 := walkAndSum(tree), findDirectoryToDelete(tree, needed)

	return part1, part2
}

func walkAndSum(dir *Dir) int {
	var sum int
	for _, n := range dir.children {
		if d, ok := n.(*Dir); ok {
			if d.Size() < 100000 {
				sum += d.Size()
			}

			sum += walkAndSum(d)
		}
	}
	return sum
}

func findDirectoryToDelete(dir *Dir, size int) int {
	// very large int
	dirSizeToDelete := math.MaxInt

	for _, n := range dir.children {
		if d, ok := n.(*Dir); ok {

			if d.Size() > size && d.Size() < dirSizeToDelete {
				dirSizeToDelete = d.Size()
			}

			innerPotential := findDirectoryToDelete(d, size)

			if innerPotential > size && innerPotential < dirSizeToDelete {
				dirSizeToDelete = innerPotential
			}
		}
	}

	return dirSizeToDelete
}

type Node interface {
	Size() int
}

type Dir struct {
	name     string
	children []Node
	parent   *Dir
}

func (d Dir) Size() int {
	var sum int
	for _, n := range d.children {
		sum += n.Size()
	}
	return sum
}

func (d Dir) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Dir %q (%d):", d.name, d.Size()))
	for _, n := range d.children {
		sb.WriteString(fmt.Sprintf("\n\t%s", n))
	}
	return sb.String()
}

type File struct {
	name string
	size int
}

func (f File) Size() int {
	return f.size
}

func (f File) String() string {
	return fmt.Sprintf("File %q (%d)", f.name, f.size)
}

func parse(s string) *Dir {
	lines := strings.Split(s, "\n")

	var (
		root       *Dir
		currentDir *Dir
	)

	for _, line := range lines {
		switch {
		case line == "$ cd /":
			// first line
			root = &Dir{name: "/"}
			currentDir = root
		case line == "$ ls":
		case strings.HasPrefix(line, "dir"):
			newDir := &Dir{name: line[4:], parent: currentDir}
			currentDir.children = append(currentDir.children, newDir)
		case regexp.MustCompile(`^\d+`).MatchString(line):
			// is file, starts with number
			s := strings.Split(line, " ")
			size, err := strconv.Atoi(s[0])
			if err != nil {
				panic(err)
			}

			file := &File{name: s[1], size: size}
			currentDir.children = append(currentDir.children, file)
		case line == "$ cd ..":
			currentDir = currentDir.parent
		case strings.HasPrefix(line, "$ cd "):
			name := line[5:]

			var dir *Dir
			for _, n := range currentDir.children {
				if d, ok := n.(*Dir); ok && d.name == name {
					dir = n.(*Dir)
					break
				}
			}

			if dir == nil {
				panic("dir not found")
			}

			currentDir = dir
		}
	}

	return root
}

func main() {
	execute.Run(run, nil, puzzle, true)
}
