package tree

import (
	"fmt"
)

type Node struct {
	Parent   *Node
	Children []*Node
	Name     string
	IsDir    bool
}

func (node *Node) Count() int {
	if node.Children == nil {
		return 0
	}

	return len(node.Children)
}

func (node *Node) Add(name string) {
	node.Children = append(
		node.Children,
		&Node{
			Parent: node,
			Name:   name,
		},
	)
}

func (node *Node) Get(name string) *Node {
	for index, child := range node.Children {
		if child.Name == name {
			return node.Children[index]
		}
	}

	return &Node{}
}

func (node *Node) Print() {
	node.print(0, "", true)
}

func (node *Node) print(level int, prefix string, isLast bool) {
	if level == 0 && prefix == "" {
		if node.Name == ".password-store" {
			fmt.Println(bold("Password Store"))
		}
	} else {
		branch := ternary(isLast, "└── ", "├── ")
		name := ternary(node.Count() > 0, bold(node.Name), node.Name)
		fmt.Println(prefix + branch + name)
		prefix += ternary(isLast, "    ", "│   ")
	}

	for index, child := range node.Children {
		child.print(level+1, prefix, index == len(node.Children)-1)
	}
}

func ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func bold(in string) string {
	return fmt.Sprintf("\033[1;37m%s\033[0m", in)
}
