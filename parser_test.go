package main

import (
	"fmt"
	"reflect"
	"testing"
)

type ParseResult[O any] struct {
	next   string
	parsed O
	err    error
}

type TestParserCore[O any] struct {
	name  string
	input string
	want  ParseResult[O]
}

func TestChar(t *testing.T) {
	tests := []struct {
		TestParserCore[string]
		target rune
	}{
		{
			TestParserCore: TestParserCore[string]{
				name:  "successful parse",
				input: "Hello world",
				want: ParseResult[string]{
					next:   "ello world",
					parsed: "H",
					err:    nil,
				},
			},
			target: 'H',
		},
		{
			TestParserCore: TestParserCore[string]{
				name:  "too short input error",
				input: "",
				want: ParseResult[string]{
					next:   "",
					parsed: "",
					err:    fmt.Errorf("input string is too short for parse"),
				},
			},
			target: 'K',
		},
		{
			TestParserCore: TestParserCore[string]{
				name:  "character does not match error",
				input: "Another message",
				want: ParseResult[string]{
					next:   "",
					parsed: "",
					err:    fmt.Errorf("character does not match"),
				},
			},
			target: 'r',
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := Char(tc.target)(tc.input)
			got := ParseResult[string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})

	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		TestParserCore[string]
		target string
	}{
		{
			TestParserCore: TestParserCore[string]{
				name:  "successful parse",
				input: "This is a message",
				want: ParseResult[string]{
					next:   " is a message",
					parsed: "This",
					err:    nil,
				},
			},
			target: "This",
		},
		{
			TestParserCore: TestParserCore[string]{
				name:  "too short input error",
				input: "short",
				want: ParseResult[string]{
					next:   "",
					parsed: "",
					err:    fmt.Errorf("input string is too short for parse"),
				},
			},
			target: "Long target",
		},
		{
			TestParserCore: TestParserCore[string]{
				name:  "target does not match error",
				input: "My test string",
				want: ParseResult[string]{
					next:   "",
					parsed: "",
					err:    fmt.Errorf("target does not match"),
				},
			},
			target: "Your",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := Match(tc.target)(tc.input)
			got := ParseResult[string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}

}

func TestTake(t *testing.T) {
	tests := []struct {
		TestParserCore[string]
		amount uint
	}{
		{
			TestParserCore: TestParserCore[string]{
				name:  "successful parse",
				input: "123456789",
				want: ParseResult[string]{
					next:   "56789",
					parsed: "1234",
					err:    nil,
				},
			},
			amount: 4,
		},
		{
			TestParserCore: TestParserCore[string]{
				name:  "too short input error",
				input: "short",
				want: ParseResult[string]{
					next:   "",
					parsed: "",
					err:    fmt.Errorf("input string is too short for parse"),
				},
			},
			amount: 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := Take(tc.amount)(tc.input)
			got := ParseResult[string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}
}

func TestOneOf(t *testing.T) {
	tests := []struct {
		TestParserCore[string]
		characters string
	}{
		{
			TestParserCore: TestParserCore[string]{
				name:  "successful parse",
				input: "abcdefg",
				want: ParseResult[string]{
					next:   "bcdefg",
					parsed: "a",
					err:    nil,
				},
			},
			characters: "xyaz",
		},
		{
			TestParserCore: TestParserCore[string]{
				name:  "too short input error",
				input: "",
				want: ParseResult[string]{
					next:   "",
					parsed: "",
					err:    fmt.Errorf("input string is too short for parse"),
				},
			},
			characters: "abc",
		},
		{
			TestParserCore: TestParserCore[string]{
				name:  "none character match error",
				input: "abc",
				want: ParseResult[string]{
					next:   "",
					parsed: "",
					err:    fmt.Errorf("none character match"),
				},
			},
			characters: "xyz",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := OneOf(tc.characters)(tc.input)
			got := ParseResult[string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}
}

func TestNoneOf(t *testing.T) {
	tests := []struct {
		TestParserCore[string]
		characters string
	}{
		{
			TestParserCore: TestParserCore[string]{
				name:  "successful parse",
				input: "abcde",
				want: ParseResult[string]{
					next:   "bcde",
					parsed: "a",
					err:    nil,
				},
			},
			characters: "xyz",
		},
		{
			TestParserCore: TestParserCore[string]{
				name:  "too short input error",
				input: "",
				want: ParseResult[string]{
					next:   "",
					parsed: "",
					err:    fmt.Errorf("input string is too short for parse"),
				},
			},
			characters: "abc",
		},
		{
			TestParserCore: TestParserCore[string]{
				name:  "some character match error",
				input: "abcde",
				want: ParseResult[string]{
					next:   "",
					parsed: "",
					err:    fmt.Errorf("some character match"),
				},
			},
			characters: "abc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := NoneOf(tc.characters)(tc.input)
			got := ParseResult[string]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTakeUntil(t *testing.T) {
	// arrange
	input := "Hello my name is AlfredoPrograma"
	separator := "name"
	expectedParsed := "Hello my "
	expectedNext := "name is AlfredoPrograma"

	// act
	next, parsed, _ := TakeUntil(separator)(input)

	// assert
	if parsed != expectedParsed {
		t.Errorf("should return the parsed character")
	}

	if next != expectedNext {
		t.Errorf("should return the rest of the input")
	}
}

func TestStrictTakeUntilSuccess(t *testing.T) {
	// arrange
	input := "Hello my name is AlfredoPrograma"
	separator := "name"
	expectedParsed := "Hello my "
	expectedNext := "name is AlfredoPrograma"

	// act
	next, parsed, _ := StrictTakeUntil(separator)(input)

	// assert
	if parsed != expectedParsed {
		t.Errorf("should return the parsed character")
	}

	if next != expectedNext {
		t.Errorf("should return the rest of the input")
	}
}

func TestStrictTakeUntilFail(t *testing.T) {
	// arrange
	input := "Hello my name is AlfredoPrograma"
	separator := "zzz"

	// act
	next, parsed, err := StrictTakeUntil(separator)(input)
	// assert
	if parsed != "" {
		t.Errorf("should return empty parsed string because it failed")
	}

	if next != "" {
		t.Errorf("should return empty rest of the input because it failed")
	}

	if err == nil {
		t.Error("should return an error because cannot match target strict parser")
	}
}
