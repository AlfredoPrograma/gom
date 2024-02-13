package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPair(t *testing.T) {
	type PairParserParams[T, K any] struct {
		first  Parser[T]
		second Parser[K]
	}

	type PairParserTestCase[T, K any] struct {
		name   string
		input  string
		params PairParserParams[T, K]
		want   ParseResult[PairResult[T, K]]
	}

	tests := []PairParserTestCase[string, string]{
		{
			name:  "successful parse",
			input: "foobarbaz",
			params: PairParserParams[string, string]{
				first:  Match("foo"),
				second: Match("bar"),
			},
			want: ParseResult[PairResult[string, string]]{
				next: "baz",
				parsed: PairResult[string, string]{
					first:  "foo",
					second: "bar",
				},
				err: nil,
			},
		},
		{
			name:  "first parser failed error",
			input: "foobarbaz",
			params: PairParserParams[string, string]{
				first:  Match("failed"),
				second: Match("bar"),
			},
			want: ParseResult[PairResult[string, string]]{
				next:   "",
				parsed: PairResult[string, string]{},
				err:    fmt.Errorf("first parser failed"),
			},
		},
		{
			name:  "second parser failed error",
			input: "foobarbaz",
			params: PairParserParams[string, string]{
				first:  Match("foo"),
				second: Match("failed"),
			},
			want: ParseResult[PairResult[string, string]]{
				next:   "",
				parsed: PairResult[string, string]{},
				err:    fmt.Errorf("second parser failed"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := Pair[string, string](tc.params.first, tc.params.second)(tc.input)
			got := ParseResult[PairResult[string, string]]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}
}
