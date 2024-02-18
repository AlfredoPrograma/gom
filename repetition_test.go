package main

import (
	"fmt"
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
