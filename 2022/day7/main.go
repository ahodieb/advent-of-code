package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ahodieb/advent-of-code-2022/common/input"
)

func main() {
	in, err := input.FromPath("2022/day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	root := Node{name: "/"}
	fs := FS{Root: &root, Curr: &root}
	digits := regexp.MustCompile(`^[0-9]+$`)

	for in.Scan() {
		cmd := strings.Split(in.Text(), " ")

		if cmd[0] == "$" {
			if cmd[1] == "cd" {
				fs.CD(cmd[2])
			}
		} else if cmd[0] == "dir" {
			fs.MKDIR(cmd[1])
		} else if digits.MatchString(cmd[0]) {
			size, _ := strconv.Atoi(cmd[0])
			fs.Touch(cmd[1], size)
		}
	}

	sum := 0
	smallest := root.Size()
	freeSpace := 70_000_000 - smallest
	neededSpace := 30_000_000 - freeSpace
	dirs := root.RecurseDirs()

	for i := range dirs {
		size := dirs[i].Size()
		if size <= 100_000 {
			sum = sum + size
		}

		if neededSpace-size <= 0 {
			if smallest > size {
				smallest = size
			}
		}
	}

	fmt.Println(sum, smallest)

}

type FS struct {
	Root *Node
	Curr *Node
}

func (f *FS) CD(d string) {
	if d == "/" {
		f.Curr = f.Root
	} else if d == ".." {
		f.Curr = f.Curr.parent
	} else {
		f.Curr = f.Curr.Find(d)
	}
}

func (f *FS) MKDIR(d string) {
	f.Curr.AddChild(&Node{name: d})
}

func (f *FS) Touch(n string, s int) {
	f.Curr.AddChild(&Node{name: n, size: s})
}

type Node struct {
	name     string
	size     int
	parent   *Node
	children []*Node
}

func (n *Node) Size() int {
	sum := 0
	for i := range n.children {
		sum += n.children[i].Size()
	}

	return n.size + sum
}

func (n *Node) AddChild(child *Node) {
	child.parent = n
	n.children = append(n.children, child)
}

func (n *Node) Find(path ...string) *Node {
	curr := n
	for _, p := range path {
		for i := range curr.children {
			if curr.children[i].name == p {
				curr = curr.children[i]
				break
			}
		}
	}

	return curr
}

func (n *Node) IsDir() bool {
	return len(n.children) != 0 && n.size == 0
}

func (n *Node) RecurseDirs() []*Node {
	nodes := []*Node{n}
	for i := range n.children {
		if !n.children[i].IsDir() {
			continue
		}

		nodes = append(nodes, n.children[i].RecurseDirs()...)
	}

	return nodes
}
