package main

import (
	"fmt"
)

type Unwrappable interface {
	unwrap(dest *any)
}

type UnwrappableString string

func (s UnwrappableString) unwrap(dest *any) {
	dest = &string(s)
}

func (s UnwrappableString) take() string {
	var dest string
	s.unwrap(&dest)

	return dest
}

type MapFn[I Unwrappable, K Unwrappable] func(parsed I) (K, error)

type Parser[O Unwrappable] func(string) (string, O, error)

/* func (p Parser[O]) Map(mapper MapFn[O, Unwrappable]) Parser[Unwrappable] {
	return func(s string) (string, string, error) {
		next, p1, _ := p(s)
		p2, err := mapper(p1)

		return next, p2, err
	}
} */

// Takes a target rune and returns a parser which matches the first rune of the input string against the target.
//
// If it matches, returns the rest of the string, the stringified rune and a nil error.
// Else returns empty values for the next string and matched string, and returns a fullfilled error.
//
// # Examples
//
//	func successExample() {
//		next, parsed, err := Char('H')("Hello world")
//		fmt.Println(next)   // "ello world"
//		fmt.Println(parsed) // "h"
//		fmt.Println(err)    // nil
//	}
//
//	func failExample() {
//		next, parsed, err := Char('K')("Hello world")
//		fmt.Println(next)   // ""
//		fmt.Println(parsed) // ""
//		fmt.Println(err)    // error
//	}
func Char(target rune) Parser[UnwrappableString] {
	return func(input string) (string, UnwrappableString, error) {
		parsed := string(input[0])

		if parsed != string(target) {
			return "", "", fmt.Errorf("character does not match")
		}

		next := input[1:]

		return next, UnwrappableString(parsed), nil
	}
}

// Takes a target string and returns a parser which matches the first set of characters of the input string against the target.
//
// If it matches, returns the rest of the string, the matched sring and a nil error.
// Else returns empty values for the next string and matched string, and returns a fullfilled error.
//
// # Examples
//
//	func successExample() {
//		next, parsed, err := Match("Hello")("Hello world")
//		fmt.Println(next)   // " world"
//		fmt.Println(parsed) // "Hello"
//		fmt.Println(err)    // nil
//	}
//
//	func failExample() {
//		next, parsed, err := Match("Invalid")("Hello world")
//		fmt.Println(next)   // ""
//		fmt.Println(parsed) // ""
//		fmt.Println(err)    // error
//	}
func Match(target string) Parser[UnwrappableString] {
	return func(input string) (string, UnwrappableString, error) {
		targetLength := len(target)

		if targetLength > len(input) {
			return "", "", fmt.Errorf("given target is longer than the input to compare")
		}

		parsed := input[:len(target)]

		if parsed != target {
			return "", "", fmt.Errorf("cannot match target against given input")
		}

		return input[len(target):], UnwrappableString(parsed), nil
	}
}
