package gofSQL

import (
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		name string
		args string
		want Queries
	}{
		{name: "Empty", args: "testfiles/test_empty.sql"},
		{name: "Complex", args: "testfiles/test_complex.sql",
			want: Queries{
				"sample":            "SELECT * FROM dummy;",
				"yet_another_query": "SELECT column_name(s) FROM table_name WHERE column_name operator ALL (SELECT column_name FROM table_name WHERE condition);",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ParseFile(tt.args); reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected: %v, got: %v", tt.want, got)
			}
		})
	}
}

var queries Queries

func benchmarkParseFile(file string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		queries, _ = ParseFile(file)
	}
}

func BenchmarkParseFile10Rows(b *testing.B) {
	benchmarkParseFile("benchfiles/bench_10-rows.sql", b)
}
func BenchmarkParseFile10000Rows(b *testing.B) {
	benchmarkParseFile("benchfiles/bench_10000-rows.sql", b)
}
func BenchmarkParseFile50000Rows(b *testing.B) {
	benchmarkParseFile("benchfiles/bench_50000-rows.sql", b)
}
func BenchmarkParseFile10000RowsSkip(b *testing.B) {
	benchmarkParseFile("benchfiles/bench_10000-rows-skip.sql", b)
}
