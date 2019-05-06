package node

import (
	h "../helper"
	"fmt"
	"os"
)

type lineType byte

const (
	ListItem         lineType = 0
	KeyValueItem     lineType = 1
	Object           lineType = 2
	Anchor           lineType = 3
	Pointer          lineType = 4
	PointerObject    lineType = 5
	ScalarTypeString lineType = 6
	ScalarTypeFloat  lineType = 7
	ScalarTypeInt    lineType = 8
	ScalarTypeBool   lineType = 9  // currently unused
	ScalarTypeNull   lineType = 10 // currently unused
	ScalarTypeBinary lineType = 11 // currently unused // represents base64 string // !!binary
	ScalarTypeDate   lineType = 12 // currently unused // represents ISO 8601 Date String
	LiteralBlock     lineType = 13 // currently unused // represents |
	FoldedBlock      lineType = 14 // currently unused // represents >
)

// TODO: Maps fixen:
// - key0: value0		(Node0)
//   key1: value1		(Node1)
// führt dazu, dass Node0 als Child Node1 hält
// Es müsste jedoch jedoch einen gemeinsamen Parent geben, der beide Objekte hält.
func getLineTypes(s string) *[]lineType {
	var types []lineType
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
	if h.Matches(s, "\\S+\\s+>\\s*$") {
		types = append(types, FoldedBlock)
	}
	if h.Matches(s, "!!\\S+") {
		_, _ = fmt.Fprintf(os.Stderr, "warning: Usage of '%v' is currently not supported\n", h.Get(s, "!!\\S+"))
	}

	return &types
}

func (l lineType) toString() string {
	switch l {
	case ListItem:
		return "ListItem"
	case KeyValueItem:
		return "KeyValueItem"
	case Object:
		return "Object"
	case Anchor:
		return "Anchor"
	case Pointer:
		return "Pointer"
	case PointerObject:
		return "PointerObject"
	case ScalarTypeString:
		return "!!str"
	case ScalarTypeFloat:
		return "!!float"
	case ScalarTypeInt:
		return "!!int"
	case ScalarTypeBool:
		return "!!bool"
	case ScalarTypeNull:
		return "!!null"
	case ScalarTypeBinary:
		return "!!binary"
	case ScalarTypeDate:
		return "!!date"
	case LiteralBlock:
		return "|"
	case FoldedBlock:
		return ">"
	default:
		return ""
	}
}
