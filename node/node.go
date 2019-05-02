package node

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	h "yaml-compare/helper"
)

const FILE = "FILE"
const BLOCK = "BLOCK"

type Node struct {
	Value    string
	parent   *Node
	children []*Node
	Indent   int
	lineType []lineType
	status   ChangeStatus
}

func NewFile() *Node {
	return &Node{Value: FILE, Indent: -2}
}
func newBlock(value string) *Node {
	return &Node{Value: BLOCK, Indent: -1, lineType: *getLineTypes(value)}
}

func New(value string) (*Node, error) {
	v := strings.Trim(h.RemoveComment(value), " ")
	if len(v) == 0 {
		return nil, errors.New("node: Comment String")
	} else if h.Matches(v, "-{3}(\\s+[>|])?") {
		return newBlock(v), nil
	} else {
		return &Node{Value: v, Indent: h.Indent(value), lineType: *getLineTypes(v)}, nil
	}
}

func (n *Node) Print() {
	fmt.Println(n.Indent, "->", n.Value, "\t\t", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(n.lineType)), ","), "[]"))
	for _, c := range n.children {
		c.Print()
	}
}

func (n *Node) PrintBy(indent int) {
	nIndent := n.Indent
	if nIndent < 0 {
		nIndent = 0
	}
	if n.Value != FILE {
		if n.status != 0 {
			fmt.Print(string(n.status))
		} else {
			fmt.Print(" ")
		}
		for i := 0; i < indent+nIndent; i++ {
			fmt.Print(" ")
		}
		if n.Value == BLOCK {
			fmt.Print("---")
		} else {
			fmt.Print(n.Value)
		}
		fmt.Println()
	}
	for _, c := range n.children {
		c.PrintBy(indent + nIndent)
	}
}

func (n *Node) AddChildren(child ...*Node) {
	n.children = append(n.children, child...)
	for _, c := range child {
		c.parent = n
	}
}

func (n *Node) GetIndentParent(node *Node) *Node {
	parent := n
	for parent.Indent >= node.Indent {
		parent = (*parent).parent
	}
	return parent
}

func (n *Node) DeleteChild(child *Node) bool {
	i := h.SliceIndex(len(n.children)-1, func(i int) bool {
		return n.children[i] == child
	})
	if i < 0 {
		return false
	}

	copy(n.children[i:], n.children[i+1:])
	n.children[len(n.children)-1] = nil // or the zero value of T
	n.children = n.children[:len(n.children)-1]

	return true
}

func (n *Node) DeleteSelf() bool {
	if n.parent == nil {
		return false
	}
	return n.parent.DeleteChild(n)
}

func (n *Node) Clean() bool {
	if strings.Compare(n.Value, FILE) == 0 && len(n.children) > 0 {
		for i := 0; i < len(n.children); {
			cleaned := n.children[i].Clean()
			if !cleaned {
				i++
			}
		}
	} else if strings.Compare(n.Value, BLOCK) == 0 && len(n.children) == 0 && n.parent != nil {
		return n.DeleteSelf()
	} else if len(n.children) > 0 {
		n.cleanChildren()
	}
	return false
}

func (n *Node) cleanChildren() {
	var r []*Node
	for _, e := range n.children {
		if e != nil {
			r = append(r, e)
		}
	}
	n.children = r
}

func (n *Node) getKey() string {
	r, _ := regexp.Compile("\\S+:")
	r0, _ := regexp.Compile(":")
	return r0.ReplaceAllString(r.FindString(n.Value), ":")
}

func (n *Node) copy() *Node {
	copied := *n
	return &copied
}

func (n *Node) deepCopy() *Node {
	copied := n.copy()

	newChilds := make([]*Node, len(n.children))
	for i, child := range n.children {
		newChilds[i] = child.deepCopy()
	}
	copied.children = newChilds
	return copied
}
