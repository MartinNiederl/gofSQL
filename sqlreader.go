package gofSQL

import (
	"bufio"
	"io"
)

type Query string
type Tag string

type Queries map[Tag]Query

func parseReader(reader io.Reader) (Queries, error) {
	var queries = make(Queries)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		ParseLine(line)
	}

	return queries, nil
}
