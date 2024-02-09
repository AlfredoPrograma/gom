package main

import "testing"

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
