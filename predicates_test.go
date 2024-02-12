package main

import (
	"fmt"
	"testing"
	"unicode"
)

func TestTakeWhile(t *testing.T) {
	tests := []ParserTestCase[Predicate, string]{
		{
			name:  "successful parse",
			input: "abcd123",
			params: func(ch rune) bool {
				return unicode.IsLetter(ch)
			},
			want: ParseResult[string]{
				next:   "123",
				parsed: "abcd",
				err:    nil,
			},
		},
		{
			name:  "predicate dont match",
			input: "1234abcd",
			params: func(ch rune) bool {
				return unicode.IsLetter(ch)
			},
			want: ParseResult[string]{
				next:   "1234abcd",
				parsed: "",
				err:    nil,
			},
		},
	}

	ExecParserTestCases[Predicate, string](t, TakeWhile, tests)
}

func TestStrictTakeWhile(t *testing.T) {
	tests := []ParserTestCase[Predicate, string]{
		{
			name:  "successful parse",
			input: "ABCDefgh",
			params: func(ch rune) bool {
				return unicode.IsUpper(ch)
			},
			want: ParseResult[string]{
				next:   "efgh",
				parsed: "ABCD",
				err:    nil,
			},
		},
		{
			name:  "predicate dont match error",
			input: "ABCDefgh",
			params: func(ch rune) bool {
				return unicode.IsLower(ch)
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("at least one character should match with the predicate"),
			},
		},
	}

	ExecParserTestCases[Predicate, string](t, StrictTakeWhile, tests)
}

func TestTakeTill(t *testing.T) {
	tests := []ParserTestCase[Predicate, string]{
		{
			name:  "successful parse",
			input: "abcd!#123",
			params: func(ch rune) bool {
				return unicode.IsDigit(ch)
			},
			want: ParseResult[string]{
				next:   "123",
				parsed: "abcd!#",
				err:    nil,
			},
		},
		{
			name:  "predicate dont match",
			input: "abc!#123",
			params: func(ch rune) bool {
				return unicode.IsLetter(ch)
			},
			want: ParseResult[string]{
				next:   "abc!#123",
				parsed: "",
				err:    nil,
			},
		},
	}

	ExecParserTestCases[Predicate, string](t, TakeTill, tests)
}

func TestStrictTakeTill(t *testing.T) {
	tests := []ParserTestCase[Predicate, string]{
		{
			name:  "successful parse",
			input: "\n\t\ta lot of spaces",
			params: func(ch rune) bool {
				return unicode.IsSpace(ch)
			},
			want: ParseResult[string]{
				next:   "a lot of spaces",
				parsed: "\n\t\t",
				err:    nil,
			},
		},
		{
			name:  "predicate dont match error",
			input: "123456789",
			params: func(ch rune) bool {
				return unicode.IsLetter(ch)
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("at least one character should match with the predicate"),
			},
		},
	}

	ExecParserTestCases[Predicate, string](t, StrictTakeTill, tests)
}
