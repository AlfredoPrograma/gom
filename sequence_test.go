package main

import (
	"fmt"
	"reflect"
	"testing"
	"unicode"
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

func TestDelimited(t *testing.T) {
	type DelimitedParserParams[T, K, O any] struct {
		opener  Parser[T]
		content Parser[O]
		closer  Parser[K]
	}

	type DelimitedParserTestCase[T, K, O any] struct {
		name   string
		input  string
		params DelimitedParserParams[T, K, O]
		want   ParseResult[O]
	}

	tests := []DelimitedParserTestCase[string, string, string]{
		{
			name:  "successful parse",
			input: "<html/> foobar",
			params: DelimitedParserParams[string, string, string]{
				opener:  Char('<'),
				content: TakeWhile(unicode.IsLetter),
				closer:  Match("/>"),
			},
			want: ParseResult[string]{
				next:   " foobar",
				parsed: "html",
				err:    nil,
			},
		},
		{
			name:  "opener parser fail error",
			input: "Hello",
			params: DelimitedParserParams[string, string, string]{
				opener:  Char('"'),
				content: TakeWhile(unicode.IsLetter),
				closer:  Char('"'),
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("opener parser failed"),
			},
		},
		{
			name:  "content parser fail error",
			input: "|123|",
			params: DelimitedParserParams[string, string, string]{
				opener:  Char('|'),
				content: StrictTakeWhile(unicode.IsLetter),
				closer:  Char('|'),
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("content parser failed"),
			},
		},
		{
			name:  "closer parser fail error",
			input: "[1,2,3}",
			params: DelimitedParserParams[string, string, string]{
				opener: Char('['),
				content: TakeWhile(func(ch rune) bool {
					return unicode.IsNumber(ch) || ch == ','
				}),
				closer: Char(']'),
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("closer parser failed"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := Delimited[string, string, string](tc.params.opener, tc.params.content, tc.params.closer)(tc.input)
			got := ParseResult[string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}
}
