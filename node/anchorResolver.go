package node

import (
	h "github.com/fluktuid/yaml-compare/helper"
	"regexp"
	"strings"
)

func (n *Node) resolveAnchors(config Config) {
	anchors, pointers, objPointers := n.getAnchorsAndPointers()
	resolve(anchors, pointers, objPointers, config)
}

func (n *Node) getAnchorsAndPointers() (*[]*Node, *[]*Node, *[]*Node) {
	var anchors []*Node
	var pointers []*Node
	var objPointers []*Node
	if contains(n.lineType, Anchor) {
		anchors = append(anchors, n)
	}
	if contains(n.lineType, Pointer) {
		pointers = append(pointers, n)
	} else if contains(n.lineType, PointerObject) {
		objPointers = append(pointers, n)
	}
	if len(n.children) > 0 {
		for _, c := range n.children {
			an, poi, objPoi := c.getAnchorsAndPointers()
			anchors = append(anchors, *an...)
			pointers = append(pointers, *poi...)
			objPointers = append(objPointers, *objPoi...)
		}
	}
	return &anchors, &pointers, &objPointers
}

func resolve(anchors *[]*Node, pointers *[]*Node, objectPointers *[]*Node, config Config) {
	anchorsMap := make(map[string]*Node)

	for _, a := range *anchors {
		anchorsMap[apValue(&a.Value)] = a
		if config.BewareAnchors {
			a.Value = h.Remove(a.Value, "\\s+$")
		} else {
			a.Value = h.Remove(a.Value, "\\*\\S+\\s+$")
		}
	}

	for _, p := range *pointers {
		apValue := apValue(&p.Value)
		anc := anchorsMap[apValue]
		if anc == nil {
			continue
		}
		if config.BewarePointer {
			p.Value = h.Remove(p.Value, "\\s+$")
		} else {
			p.Value = h.Remove(p.Value, "\\*\\S+\\s+$")
		}
		p.children = append(p.children, anc.children...)
	}

	for _, p := range *objectPointers {
		apValue := apValue(&p.Value)
		anc := anchorsMap[apValue]
		if anc == nil {
			continue
		}
		parent := p.parent
		if !config.BewarePointer {
			p.DeleteSelf()
		}
		var c []*Node
		for _, v := range anc.children {
			index := h.SliceIndex(len(parent.children), func(i int) bool {
				return strings.Compare(parent.children[i].getKey(), v.getKey()) == 0
			})
			if index < 0 {
				c = append(c, v)
			}
		}
		parent.AddChildren(c...)
		if config.BewarePointer {
			p.Value = h.Remove(p.Value, "\\s+$")
		} else {
			p.Value = h.Remove(p.Value, "\\*\\S+\\s+$")
		}
	}
}

func contains(s []lineType, e lineType) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func apValue(s *string) string {
	r, _ := regexp.Compile("[&*]\\S+")
	r0, _ := regexp.Compile("[&*]")
	return r0.ReplaceAllString(r.FindString(*s), "")
}

func remove(arr *[]*Node, i int) {
	copy((*arr)[i:], (*arr)[i+1:])
	(*arr)[len(*arr)-1] = nil // or the zero value of T
	y := (*arr)[:len(*arr)-1]
	arr = &y
}
