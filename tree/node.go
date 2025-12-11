package tree

import (
	"fmt"
)

type Node struct {
	Children []*Node
	Name     string
	Path     string
	IsDir    bool
}

func (node *Node) Count() int {
	if node.Children == nil {
		return 0
	}

	return len(node.Children)
}

func (node *Node) Print() {
	node.print(0, "", true)
}

func (node *Node) print(level int, prefix string, isLast bool) {
	if level == 0 && prefix == "" {
		fmt.Println(bold("Password Store"))
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
