package main

import (
	"testing"
	"unicode"
)

func TestFlexTakeWhile(t *testing.T) {
	// arrange
	input := "abcd123"
	expectedParsed := "abcd"
	expectedNext := "123"

	// act
	next, parsed, _ := FlexTakeWhile(func(ch rune) bool {
		return unicode.IsLetter(ch)
	})(input)

	// assert
	if parsed != expectedParsed {
		t.Errorf("should return the parsed character")
	}

	if next != expectedNext {
		t.Errorf("should return the rest of the input")
	}
}
