package gom

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
			next, parsed, err := Pair(tc.params.first, tc.params.second)(tc.input)
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
			next, parsed, err := Delimited(tc.params.opener, tc.params.content, tc.params.closer)(tc.input)
			got := ParseResult[string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}
}

func TestPreceded(t *testing.T) {
	type PrecededParserParams[T, O any] struct {
		preceded Parser[T]
		content  Parser[O]
	}

	type PrecededParserTestCase[T, O any] struct {
		name   string
		input  string
		params PrecededParserParams[T, O]
		want   ParseResult[string]
	}

	tests := []PrecededParserTestCase[string, string]{
		{
			name:  "successful parse",
			input: "prefixHelloSuffix",
			params: PrecededParserParams[string, string]{
				preceded: Match("prefix"),
				content:  Match("Hello"),
			},
			want: ParseResult[string]{
				next:   "Suffix",
				parsed: "Hello",
				err:    nil,
			},
		},
		{
			name:  "preceded parser fail error",
			input: "<tag/>",
			params: PrecededParserParams[string, string]{
				preceded: Match("{"),
				content:  Match("tag/>"),
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("preceded parser failed"),
			},
		},
		{
			name:  "content parser fail error",
			input: "[PARSER",
			params: PrecededParserParams[string, string]{
				preceded: Char('['),
				content:  StrictTakeWhile(unicode.IsLower),
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("content parser failed"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := Preceded(tc.params.preceded, tc.params.content)(tc.input)
			got := ParseResult[string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTerminated(t *testing.T) {
	type TerminatedParserParams[T, O any] struct {
		terminated Parser[T]
		content    Parser[O]
	}

	type TerminatedParserTestCase[T, O any] struct {
		name   string
		input  string
		params TerminatedParserParams[T, O]
		want   ParseResult[string]
	}

	tests := []TerminatedParserTestCase[string, string]{
		{
			name:  "successful parse",
			input: "hello>>>",
			params: TerminatedParserParams[string, string]{
				content:    Match("hello"),
				terminated: Match(">>>"),
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "hello",
				err:    nil,
			},
		},
		{
			name:  "content parser fail error",
			input: "12345>",
			params: TerminatedParserParams[string, string]{
				content:    StrictTakeWhile(unicode.IsLetter),
				terminated: Char('>'),
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("content parser failed"),
			},
		},
		{
			name:  "terminated parser fail error",
			input: "abcde<<<",
			params: TerminatedParserParams[string, string]{
				content:    StrictTakeWhile(unicode.IsLetter),
				terminated: StrictTakeWhile(unicode.IsDigit),
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("terminated parser failed"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := Terminated(tc.params.content, tc.params.terminated)(tc.input)
			got := ParseResult[string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}
}
