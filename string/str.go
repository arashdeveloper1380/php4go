package string

import (
	"strconv"
	"strings"
)

type Str struct{}

func StrInstance() Str {
	return Str{}
}

func (str *Str) StrLen(value string) (string, error) {
	length := len(value)
	lengthStr := strconv.Itoa(length)
	return lengthStr, nil
}

// find index string
func (str *Str) StrPos(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// StrRPos finds the last occurrence of the needle in the haystack.
func (str *Str) StrRPos(haystack string, needle string) int {
	return strings.LastIndex(haystack, needle)
}

func (str *Str) StrReplace(haystack string old, new string) string{
	return strings.ReplaceAll(haystack, old, new)
}

func (str *Str) SubStr(s string, start, length int) string {
	if start < 0 {
		start = 0
	}
	if start >= len(s) {
		return ""
	}
	end := start + length
	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}

// TrimSpace removes spaces
func (str *Str) TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// Trim removes the specified cutset of characters from the beginning and end of the input string.
func (str *Str) Trim(s, cutset string) string {
	return strings.Trim(s, cutset)
}

// TrimLeft removes leading characters specified in the cutset from the input string.
func (str *Str) TrimLeft(s, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// TrimRight removes trailing characters specified in the cutset from the input string.
func (str *Str) TrimRight(s, cutset string) string {
	return strings.TrimRight(s, cutset)
}