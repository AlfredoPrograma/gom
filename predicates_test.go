package main

import (
	"testing"
	"unicode"
)

func TestTakeWhile(t *testing.T) {
	// arrange
	input := "abcd123"
	expectedParsed := "abcd"
	expectedNext := "123"

	// act
	next, parsed, _ := TakeWhile(func(ch rune) bool {
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

func TestStrictTakeWhileSuccess(t *testing.T) {
	// arrange
	input := "abcd123"
	expectedParsed := "abcd"
	expectedNext := "123"

	// act
	next, parsed, _ := StrictTakeWhile(func(ch rune) bool {
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

func TestStrictTakeWhileFail(t *testing.T) {
	// arrange
	input := "abcd123"

	// act
	next, parsed, err := StrictTakeWhile(func(ch rune) bool {
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
