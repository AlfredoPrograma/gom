package gom

import (
	"fmt"
	"reflect"
	"testing"
	"unicode"
)

func TestMany(t *testing.T) {
	tests := []ParserTestCase[Parser[string], []string]{
		{
			name:   "successful parse",
			input:  "ababab123",
			params: Match("ab"),
			want: ParseResult[[]string]{
				next: "123",
				parsed: []string{
					"ab",
					"ab",
					"ab",
				},
				err: nil,
			},
		},
		{
			name:   "successful empty parse",
			input:  "ababab123",
			params: Match("xd"),
			want: ParseResult[[]string]{
				next:   "ababab123",
				parsed: []string{},
				err:    nil,
			},
		},
	}

	ExecParserTestCases[Parser[string], []string](t, Many, tests)
}

func TestStrictMany(t *testing.T) {
	tests := []ParserTestCase[Parser[string], []string]{
		{
			name:   "successful parse",
			input:  "    four whitespaces",
			params: Char(' '),
			want: ParseResult[[]string]{
				next:   "four whitespaces",
				parsed: []string{" ", " ", " ", " "},
				err:    nil,
			},
		},
		{
			name:   "parser dont match at leas one time",
			input:  ">>>Hello<<<",
			params: StrictTakeWhile(unicode.IsNumber),
			want: ParseResult[[]string]{
				next:   "",
				parsed: []string{},
				err:    fmt.Errorf("parser should match at least one time"),
			},
		},
	}

	ExecParserTestCases[Parser[string], []string](t, StrictMany, tests)
}

func TestCount(t *testing.T) {
	type CountParserParams[O any] struct {
		times  uint
		parser Parser[O]
	}

	type CountParserTestCase[O any] struct {
		name   string
		input  string
		params CountParserParams[O]
		want   ParseResult[[]O]
	}

	tests := []CountParserTestCase[string]{
		{
			name:  "successful parse",
			input: "123456end",
			params: CountParserParams[string]{
				parser: Take(2),
				times:  3,
			},
			want: ParseResult[[]string]{
				next:   "end",
				parsed: []string{"12", "34", "56"},
				err:    nil,
			},
		},
		{
			name:  "parser fail error",
			input: "notAdot",
			params: CountParserParams[string]{
				parser: Char('.'),
				times:  2,
			},
			want: ParseResult[[]string]{
				next:   "",
				parsed: []string{},
				err:    fmt.Errorf("cannot execute parser %v times", 2),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := Count[string](tc.params.parser, tc.params.times)(tc.input)
			got := ParseResult[[]string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}
}
