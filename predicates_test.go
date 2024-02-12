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
	// arrange
	input := "abc!#123"
	expectedParsed := "abc!#"
	expectedNext := "123"

	// act
	next, parsed, _ := TakeTill(func(ch rune) bool {
		return unicode.IsDigit(ch)
	})(input)

	// assert
	if parsed != expectedParsed {
		t.Errorf("should return the parsed character")
	}

	if next != expectedNext {
		t.Errorf("should return the rest of the input")
	}
}

func TestStrictTakeTillFail(t *testing.T) {
	// arrange
	input := "abcd123"

	// act
	next, parsed, err := StrictTakeTill(func(ch rune) bool {
		return unicode.IsDigit(ch)
	})(input)

	// assert
	if parsed != "" {
		t.Errorf("should return empty parsed string because it failed")
	}

	if next != "" {
		t.Errorf("should return empty rest of the input because it failed")
	}

	if err == nil {
		t.Error("should return an error because none character complied predicate on strict parser")
	}
}
