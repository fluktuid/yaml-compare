package node

import (
	"fmt"
	"regexp"
)

type Node struct {
	Value    string
	parent   *Node
	children []*Node
	Indent   int
}

func NewNode(value string) *Node {
	return &Node{Value: value, Indent: getIndent(value)}
}

func (n Node) ToString() {
	fmt.Println(n.Indent, "->", n.Value)
	for _, c := range n.children {
		c.ToString()
	}
}

func (n *Node) AddChildren(child *Node) {
	n.children = append(n.children, child)
	child.parent = n
}

func (n *Node) GetIndentParent(node *Node) *Node {
	parent := n
	for parent.Indent >= node.Indent {
		parent = (*parent).parent
	}
	return parent
}

func getIndent(line string) int {
	r, _ := regexp.Compile("^\\s*")
	return len(r.FindString(line))
}
