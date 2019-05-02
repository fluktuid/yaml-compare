package node

import (
	"reflect"
	h "yaml-compare/helper"
)

func (n *Node) Compare(other *Node) *Node {
	return n.getDifference(other)
}

func (n *Node) getDifference(other *Node) *Node {
	if n.same(other) {
		return nil
	}
	difference := n.deepCopy()

	thisSpare := n.children
	otherSpare := other.children

	var joined []Pair

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
			if indexOfEqual < 0 {
				continue
			}
		}
		joined = append(joined, Pair{first: v, second: otherSpare[indexOfEqual]})
		remove(&otherSpare, indexOfEqual)
	}

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
			if indexOfEqual < 0 {
				continue
			}
		}
		joined = append(joined, Pair{first: v, second: otherSpare[indexOfEqual]})
		remove(&thisSpare, indexOfEqual)
	}

	var differentChilds []*Node
	for _, v := range joined {
		if v.second != nil {
			differentChilds = append(differentChilds, v.first.getDifference(v.second))
		} else {
			differentChilds = append(differentChilds, v.first)
		}
	}

	/*
	   if (same(other))
	       return null

	   val y = clone()

	   val thisSpare = mutableListOf(* children.toTypedArray())
	   val otherSpare = mutableListOf(* other.children.toTypedArray())

	   val joined = thisSpare.map {
	       val sameFromOther =
	           otherSpare.firstOrNull { o -> o.same(it) } ?: otherSpare.firstOrNull { o -> o.value == it.value }
	       otherSpare.remove(sameFromOther)
	       Pair(it, sameFromOther)
	   }.toMutableList()
	   joined.addAll(
	       otherSpare.map {
	           val sameFromThis =
	               thisSpare.firstOrNull { o -> o.same(it) } ?: thisSpare.firstOrNull { o -> o.value == it.value }
	           thisSpare.remove(sameFromThis)
	           Pair(it, sameFromThis)
	       }
	   )

	   val differentChilds = joined.mapNotNull {
	       when {
	           it.second != null -> {
	               it.first.getDifference(it.second!!)
	           }
	           else -> it.first
	       }
	   }

	   y.addChild(differentChilds)
	   return y
	*/
	difference.children = differentChilds
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
	if n.parent != other.parent {
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
