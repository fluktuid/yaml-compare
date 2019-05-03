package helper

import (
	"regexp"
	"strconv"
)

func RemoveComment(s string) string {
	r, _ := regexp.Compile("#")
	matchIndex := r.FindStringIndex(s)
	if matchIndex == nil {
		return s
	}
	return s[:matchIndex[0]]
}

func Trim(s string) string {
	r, _ := regexp.Compile("(^\\s+|\\s+$)")
	return r.ReplaceAllString(s, "")
}

func TrimExt(s string) string {
	r, _ := regexp.Compile("(^\\s+|\\s+$|:|-\\s+)")
	return r.ReplaceAllString(s, "")
}

func MapListString(s string, position int) *string {
	// full key match: "^(\\s*\\-\\s+)?(\\S[^\\.]+)"
	r, _ := regexp.Compile("(\\S[^.]+)")
	var pos = r.FindStringIndex(s)
	if pos != nil {
		s = s[:pos[1]]
	}
	if position < 0 {
		return &s
	}
	s = s + "[" + strconv.Itoa(position) + "]"
	return &s
}

func Remove(s string, regex string) string {
	r, e := regexp.Compile(regex)
	if e != nil {
		return s
	}
	return r.ReplaceAllString(s, "")
}

func Matches(s string, regex string) bool {
	r, e := regexp.Compile(regex)
	if e != nil {
		return false
	}
	return r.MatchString(s)
}

func Get(s string, regex string) string {
	r, e := regexp.Compile(regex)
	if e != nil {
		return ""
	}
	return r.FindString(s)
}

func Indent(line string) int {
	r, _ := regexp.Compile("^\\s*")
	return len(r.FindString(line))
}
