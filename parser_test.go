package main

import (
	"fmt"
	"testing"
)

func TestChar(t *testing.T) {
	tests := []ParserTestCase[rune, string]{
		{
			name:   "successful parse",
			input:  "Hello world",
			params: 'H',
			want: ParseResult[string]{
				next:   "ello world",
				parsed: "H",
				err:    nil,
			},
		},
		{
			name:   "too short input error",
			input:  "",
			params: 'K',
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("input string is too short for parse"),
			},
		},
		{
			name:   "character does not match error",
			input:  "Another message",
			params: 'r',
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("character does not match"),
			},
		},
	}

	ExecParserTestCases[rune, string](t, Char, tests)
}

func TestMatch(t *testing.T) {
	tests := []ParserTestCase[string, string]{
		{
			name:   "successful parse",
			input:  "This is a message",
			params: "This",
			want: ParseResult[string]{
				next:   " is a message",
				parsed: "This",
				err:    nil,
			},
		},
		{
			name:   "too short input error",
			input:  "short",
			params: "Long params",
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("input string is too short for parse"),
			},
		},
		{
			name:   "params does not match error",
			input:  "My test string",
			params: "Your",
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("target does not match"),
			},
		},
	}

	ExecParserTestCases[string, string](t, Match, tests)
}

func TestTake(t *testing.T) {
	tests := []ParserTestCase[uint, string]{
		{
			name:   "successful parse",
			input:  "123456789",
			params: 4,
			want: ParseResult[string]{
				next:   "56789",
				parsed: "1234",
				err:    nil,
			},
		},
		{
			name:   "too short input error",
			input:  "short",
			params: 10,
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("input string is too short for parse"),
			},
		},
	}

	ExecParserTestCases[uint, string](t, Take, tests)

}

func TestOneOf(t *testing.T) {
	tests := []ParserTestCase[string, string]{
		{
			name:   "successful parse",
			input:  "abcdefg",
			params: "xyaz",
			want: ParseResult[string]{
				next:   "bcdefg",
				parsed: "a",
				err:    nil,
			},
		},
		{
			name:   "too short input error",
			input:  "",
			params: "abc",
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("input string is too short for parse"),
			},
		},
		{
			name:   "none character match error",
			input:  "abc",
			params: "xyz",
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("none character match"),
			},
		},
	}

	ExecParserTestCases[string, string](t, OneOf, tests)
}

func TestNoneOf(t *testing.T) {
	tests := []ParserTestCase[string, string]{
		{
			name:   "successful parse",
			input:  "abcde",
			params: "xyz",
			want: ParseResult[string]{
				next:   "bcde",
				parsed: "a",
				err:    nil,
			},
		},
		{
			name:   "too short input error",
			input:  "",
			params: "abc",
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("input string is too short for parse"),
			},
		},
		{
			name:   "some character match error",
			input:  "abcde",
			params: "abc",
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("some character match"),
			},
		},
	}

	ExecParserTestCases[string, string](t, NoneOf, tests)
}

func TestTakeUntil(t *testing.T) {
	tests := []ParserTestCase[string, string]{
		{
			name:   "successful parse",
			input:  "Hello my name is FooBar",
			params: "name",
			want: ParseResult[string]{
				next:   "name is FooBar",
				parsed: "Hello my ",
				err:    nil,
			},
		},
		{
			name:   "params dont match",
			input:  "Hello my name is FooBar",
			params: "people",
			want: ParseResult[string]{
				next:   "Hello my name is FooBar",
				parsed: "",
				err:    nil,
			},
		},
	}

	ExecParserTestCases[string, string](t, TakeUntil, tests)
}

func TestStrictTakeUntil(t *testing.T) {
	tests := []ParserTestCase[string, string]{
		{
			name:   "successful parse",
			input:  "Hello my name is StrictFooBar",
			params: "name",
			want: ParseResult[string]{
				next:   "name is StrictFooBar",
				parsed: "Hello my ",
				err:    nil,
			},
		},
		{
			name:   "params dont match error",
			input:  "Hello my name is StrictFooBar",
			params: "people",
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("cannot match target"),
			},
		},
	}

	ExecParserTestCases[string, string](t, StrictTakeUntil, tests)
}
