package node

import (
	"reflect"
	"regexp"
	h "yaml-compare/helper"
)

func (n *Node) Compare(other *Node) *Node {
	return n.getDifference(other)
}

func (n *Node) getDifference(other *Node) *Node {
	if n.same(other) {
		return nil
	}
	difference := n.copy()
	difference.children = nil

	thisSpare := make([]*Node, len(n.children))
	copy(thisSpare, n.children)
	otherSpare := make([]*Node, len(other.children))
	copy(otherSpare, other.children)

	var joined []Pair

	// check for changed
	for i, v := range otherSpare {
		if v == nil || len(v.children) > 0 {
			continue
		}

		r, _ := regexp.Compile("\\:(.)*$") //Compile("^\\s*\\-?\\s*\\S+\\:\\s+")

		vVal := r.ReplaceAllString(v.Value, "")
		if len(vVal) == 0 {
			continue
		}
		indexOfEqual := h.SliceIndex(len(thisSpare), func(i int) bool {
			if thisSpare[i] == nil {
				return false
			}
			tVal := r.ReplaceAllString(thisSpare[i].Value, "")
			return tVal == vVal && thisSpare[i].Value != v.Value
		})
		if indexOfEqual < 0 {
			continue
		}
		v.status = CHANGED
		remove(&thisSpare, indexOfEqual)
		remove(&otherSpare, i)
		joined = append(joined, Pair{first: v, second: nil})
	}

	// check for added
	for _, v := range otherSpare {
		if v == nil {
			continue
		}
		indexOfEqual := h.SliceIndex(len(thisSpare), func(i int) bool {
			if thisSpare[i] == nil {
				return false
			}
			return reflect.DeepEqual(*thisSpare[i], *v)
		})
		if indexOfEqual < 0 {
			indexOfEqual = h.SliceIndex(len(thisSpare), func(i int) bool {
				if thisSpare[i] == nil {
					return false
				}
				return thisSpare[i].Value == v.Value
			})
		}
		if indexOfEqual < 0 {
			// added
			joined = append(joined, Pair{first: nil, second: v})
		} else {
			joined = append(joined, Pair{first: v, second: thisSpare[indexOfEqual]})
			remove(&thisSpare, indexOfEqual)
		}
	}

	// check for removed
	for _, v := range thisSpare {
		if v == nil {
			continue
		}
		indexOfEqual := h.SliceIndex(len(otherSpare), func(i int) bool {
			if otherSpare[i] == nil {
				return false
			}
			return reflect.DeepEqual(*otherSpare[i], *v)
		})
		if indexOfEqual < 0 {
			indexOfEqual = h.SliceIndex(len(otherSpare), func(i int) bool {
				if otherSpare[i] == nil {
					return false
				}
				return otherSpare[i].Value == v.Value
			})
		}
		if indexOfEqual < 0 {
			// removed
			joined = append(joined, Pair{first: v, second: nil})
		} else {
			joined = append(joined, Pair{first: v, second: otherSpare[indexOfEqual]})
			remove(&otherSpare, indexOfEqual)
		}
	}

	var differentChilds []*Node
	for _, v := range joined {
		if v.first != nil && v.second != nil {
			change := append(differentChilds, v.first.getDifference(v.second))
			if change != nil {
				differentChilds = change
			}
		} else if v.first != nil {
			if v.first.status != CHANGED {
				v.first.status = REMOVED
			}
			differentChilds = append(differentChilds, v.first)
		} else if v.second != nil {
			v.second.status = ADDED
			differentChilds = append(differentChilds, v.second)
		}
	}

	difference.children = differentChilds
	difference.cleanChildren()
	difference.status = UNDEFINED
	return difference
}

type Pair struct {
	first  *Node
	second *Node
}

func (n *Node) same(other *Node) bool {
	if n.Value != other.Value {
		return false
	}
	if n.Indent != other.Indent {
		return false
	}
	if len(n.children) != len(other.children) {
		return false
	}

	for i := range n.children {
		if &n.children[i] != &other.children[i] {
			return false
		}
	}
	return true
}
