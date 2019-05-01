package node

import (
	"regexp"
	h "yaml-compare/helper"
)

func (n *Node) ResolveAnchors() {
	anchors, pointers := n.getAnchorsAndPointers()
	resolve(anchors, pointers)
}

func (n *Node) getAnchorsAndPointers() (*[]*Node, *[]*Node) {
	var anchors []*Node
	var pointers []*Node
	if contains(n.lineType, Anchor) {
		anchors = append(anchors, n)
	}
	if contains(n.lineType, Pointer) {
		pointers = append(pointers, n)
	} else if contains(n.lineType, PointerObject) {
		pointers = append(pointers, n)
	}
	if len(n.children) > 0 {
		for _, c := range n.children {
			an, poi := c.getAnchorsAndPointers()
			anchors = append(anchors, *an...)
			pointers = append(pointers, *poi...)
		}
	}
	return &anchors, &pointers
}

func resolve(anchors *[]*Node, pointers *[]*Node) {
	anchorsMap := make(map[string]*Node)

	for _, a := range *anchors {
		anchorsMap[apValue(&a.Value)] = a
	}

	for _, p := range *pointers {
		apValue := apValue(&p.Value)
		anc := anchorsMap[apValue]
		if anc == nil {
			continue
		}
		p.Value = h.Remove(p.Value, "[&*]\\S+")
		p.children = append(p.children, anc.children...)
	}
}

func contains(s []LineType, e LineType) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/*

fun String.getAPValue() = Regex("[&*]\\S+").find(this)?.value?.replace(Regex("[&*]"), "")
*/
func apValue(s *string) string {
	r, _ := regexp.Compile("[&*]\\S+")
	r0, _ := regexp.Compile("[&*]")
	return r0.ReplaceAllString(r.FindString(*s), "")
}
