package main

import (
	"fmt"
)

type Parser[O any] func(string) (string, O, error)

// Takes a target rune and returns a parser which matches the first rune of the input string against the target.
//
// If it matches, returns the rest of the string, the matched rune and a nil error.
// Else returns empty values for the string and rune, and returns a fullfilled error.
//
// # Examples
//
//	func successExample() {
//		next, parsed, err := Char('H')("Hello world")
//		fmt.Println(next)   // "ello world"
//		fmt.Println(parsed) // 'h'
//		fmt.Println(err)    // nil
//	}
//
//	func failExample() {
//		next, parsed, err := Char('K')("Hello world")
//		fmt.Println(next)   // ""
//		fmt.Println(parsed) // 0
//		fmt.Println(err)    // error
//	}
func Char(target rune) Parser[string] {
	return func(input string) (string, string, error) {
		parsed := string(input[0])

		if parsed != string(target) {
			return "", "", fmt.Errorf("character does not match")
		}

		next := input[1:]

		return next, parsed, nil
	}
}
