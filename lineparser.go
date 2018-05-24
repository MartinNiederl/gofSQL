package gofSQL

import (
	"regexp"
	"strings"
)

const (
	lineEmpty = iota
	lineTag
	lineQuery
	lineComment
	lineIgnore
)

var (
	regexTag     = regexp.MustCompile("^-- name: \\w+$")
	regexComment = regexp.MustCompile("^.*--.*$")
	regexIgnore  = regexp.MustCompile("^-- ignore$")
)

type ParsedLine struct {
	Type  int
	Value string
}

func ParseLine(line string) ParsedLine {
	line = strings.TrimSpace(line)

	var pLine = ParsedLine{
		Type:  lineEmpty,
		Value: line,
	}

	if line == "" {
		return pLine
	}

	if regexIgnore.MatchString(line) {
		pLine.Type = lineIgnore
		return pLine
	}

	if regexTag.MatchString(line) {
		pLine.Type = lineTag
		pLine.Value = line[len("-- name: "):]
		return pLine
	}

	if regexComment.MatchString(line) {
		pLine.Type = lineComment
		return pLine
	}

	pLine.Type = lineQuery
	return pLine
}
