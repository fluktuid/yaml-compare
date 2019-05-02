package node

import (
	h "yaml-compare/helper"
)

type LineType int

const (
	ListItem         LineType = 0
	KeyValueItem     LineType = 1
	Object           LineType = 2
	Anchor           LineType = 3
	Pointer          LineType = 4
	PointerObject    LineType = 5
	ScalarTypeString LineType = 6
	ScalarTypeFloat  LineType = 7
	ScalarTypeInt    LineType = 8
	ScalarTypeBool   LineType = 9  // currently unused
	ScalarTypeNull   LineType = 10 // currently unused
	ScalarTypeBinary LineType = 11 // currently unused // represents base64 string // !!binary
	ScalarTypeDate   LineType = 12 // currently unused // represents ISO 8601 Date String
	LiteralBlock     LineType = 13 // currently unused // represents |
	FoldedBlock      LineType = 14 // currently unused // represents >
)

// TODO: Maps fixen:
// - key0: value0		(Node0)
//   key1: value1		(Node1)
// führt dazu, dass Node0 als Child Node1 hält
// Es müsste jedoch jedoch einen gemeinsamen Parent geben, der beide Objekte hält.
func getLineTypes(s string) *[]LineType {
	var types []LineType
	if h.Matches(s, "^\\s*-\\s+\\S+") {
		types = append(types, ListItem)
	}
	if h.Matches(s, "\\S+:\\s+[^&*]*$") {
		types = append(types, KeyValueItem)
	}
	if h.Matches(s, "\\S+:\\s*$") {
		types = append(types, Object)
	}
	if h.Matches(s, "&\\S+") {
		types = append(types, Anchor)
	}
	if h.Matches(s, "\\s*<<:\\s+\\S+\\s*$") {
		types = append(types, PointerObject)
	} else if h.Matches(s, "\\*\\S+") {
		types = append(types, Pointer)
	}
	if h.Matches(s, "\\S+\\s+\\|\\s*$") {
		types = append(types, LiteralBlock)
	}
	if h.Matches(s, "\\S+\\s+>\\s*$") {
		types = append(types, FoldedBlock)
	}
	return &types
}
