package gofSQL

import "testing"

func TestParseLine(t *testing.T) {
	tests := []struct {
		name string
		args string
		want ParsedLine
	}{
		{name: "Tag", args: "-- name: sim_ple21te_st ", want: ParsedLine{Type: lineTag, Value: "sim_ple21te_st"}},
		{name: "Query", args: "SELECT * FROM simple;", want: ParsedLine{Type: lineQuery, Value: "SELECT * FROM simple;"}},
		{name: "Comment", args: "-- just  a  comment :)", want: ParsedLine{Type: lineComment, Value: "-- just  a  comment :)"}},
		{name: "Empty", args: " ", want: ParsedLine{Type: lineEmpty, Value: ""}},
		{name: "Ignore", args: "-- ignore", want: ParsedLine{Type: lineIgnore, Value: ""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLine(tt.args); got != tt.want {
				t.Errorf("expected: %v, got: %v", tt.want, got)
			}
		})
	}
}

var parsedLine ParsedLine

func benchmarkParseLine(line string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		parsedLine = ParseLine(line)
	}
}

func BenchmarkParseLineTag(b *testing.B)     { benchmarkParseLine("-- name: sim_ple21te_st ", b) }
func BenchmarkParseLineQuery(b *testing.B)   { benchmarkParseLine("SELECT * FROM simple;", b) }
func BenchmarkParseLineComment(b *testing.B) { benchmarkParseLine("-- just  a  comment :)", b) }
func BenchmarkParseLineEmpty(b *testing.B)   { benchmarkParseLine(" ", b) }
func BenchmarkParseLineIgnore(b *testing.B)  { benchmarkParseLine("-- ignore", b) }
func BenchmarkParseLineLongerQuery(b *testing.B) {
	benchmarkParseLine("SELECT EMP_ID, LAST_NAME FROM EMPLOYEE_TBL WHERE CITY = 'INDIANAPOLIS' ORDER BY 2, 1;", b)
}
