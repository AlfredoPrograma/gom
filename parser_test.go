package main

import (
	"testing"
)

func TestChar(t *testing.T) {
	// arrange
	input := "Hello world"
	expectedParsed := string(input[0])
	expectedNext := input[1:]

	// act
	next, parsed, _ := Char('H')(input)

	// assert
	if parsed != expectedParsed {
		t.Errorf("should return the parsed character")
	}

	if next != expectedNext {
		t.Errorf("should return the rest of the input")
	}
}

func TestMatch(t *testing.T) {
	// arrange
	input := "Hello world"
	expectedParsed := "Hello"
	expectedNext := " world"

	// act
	next, parsed, _ := Match("Hello")(input)

	// assert
	if parsed != expectedParsed {
		t.Errorf("should return parsed substring")
	}

	if next != expectedNext {
		t.Errorf("should return the rest of the input")
	}
}

func TestTake(t *testing.T) {
	// arrange
	input := "123456789"
	expectedParsed := "123"
	expectedNext := "456789"

	// act
	next, parsed, _ := Take(3)(input)

	// assert
	if parsed != expectedParsed {
		t.Errorf("should return parsed substring taken by the given number of characters requested")
	}

	if next != expectedNext {
		t.Errorf("should return the rest of the input")
	}
}

func TestOneOf(t *testing.T) {
	// arrange
	input := "abcde"
	expectedParsed := "a"
	expectedNext := "bcde"

	// act
	next, parsed, _ := OneOf("cbax")(input)

	// assert
	if parsed != expectedParsed {
		t.Errorf("should return the parsed character")
	}

	if next != expectedNext {
		t.Errorf("should return the rest of the input")
	}
}

func TestNoneOf(t *testing.T) {
	// arrange
	input := "abcde"
	expectedParsed := "a"
	expectedNext := "bcde"

	// act
	next, parsed, _ := NoneOf("xyz")(input)

	// assert
	if parsed != expectedParsed {
		t.Errorf("should return the parsed character")
	}

	if next != expectedNext {
		t.Errorf("should return the rest of the input")
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
