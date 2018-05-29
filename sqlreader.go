package gofSQL

import (
	"bufio"
	"bytes"
	"io"
)

const (
	typeCreate QueryType = iota
	typeInsert
	typeQuery
	typePrepare
	typeUnset
)

type QueryType int

type Query struct {
	Value string
	QType QueryType
}

type Tag string

type Queries map[Tag]Query

func parseReader(reader io.Reader) (Queries, error) {
	var (
		queries      = make(Queries)
		currTag  Tag = ""
		skipNext     = false
	)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := ParseLine(line)

		if skipNext && parsedLine.Type != lineTag {
			continue
		}

		switch parsedLine.Type {
		case lineEmpty, lineComment:
			continue
		case lineIgnore:
			skipNext = true
		case lineTag:
			skipNext = false
			currTag = Tag(parsedLine.Value)
		case lineQuery:
			query := getQuery(parsedLine.Value, scanner)
			if query != "" {
				queries[currTag] = Query{
					Value: query,
					QType: typeUnset,
				}
			}
		}
	}

	return queries, nil
}

func getQuery(part string, scanner *bufio.Scanner) string {
	var queryBuffer bytes.Buffer
	queryBuffer.WriteString(part + " ")

	for scanner.Scan() {
		parsedLine := ParseLine(scanner.Text())

		if parsedLine.Type == lineQuery {
			queryBuffer.WriteString(parsedLine.Value)
		} else {
			break
		}
	}

	return queryBuffer.String()
}
