package helper

import "regexp"

func RemoveComment(s string) string {
	r, _ := regexp.Compile("#")
	matchIndex := r.FindStringIndex(s)
	if matchIndex == nil {
		return s
	}
	return s[:matchIndex[0]]
}

func Trim(s string) string {
	r, _ := regexp.Compile("^\\s+")
	s = r.ReplaceAllString(s, "")
	r, _ = regexp.Compile("\\s+$")
	return r.ReplaceAllString(s, "")
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
