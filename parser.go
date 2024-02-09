package main

import (
	"fmt"
)

// Describes the signature function for parsers.
type Parser[O any] func(string) (string, O, error)

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

// Takes a target string and returns a parser which matches the first set of characters of the input string against the target.
//
// If it matches, returns the rest of the string, the matched string and a nil error.
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
func Match(target string) Parser[string] {
	return func(input string) (string, string, error) {
		targetLength := len(target)

		if targetLength > len(input) {
			return "", "", fmt.Errorf("given target is longer than the input to compare")
		}

		parsed := input[:len(target)]

		if parsed != target {
			return "", "", fmt.Errorf("cannot match target against given input")
		}

		return input[len(target):], parsed, nil
	}
}

// Takes an unsigned integer and returns a parser which takes the first n characters from the beginning of the input string.
//
// If it matches, returns the rest of the string, string built from the taken characters and a nil error.
// Else returns empty values for the next string and matched string, and returns a fullfilled error.
//
// # Examples
//
//	func successExample() {
//		next, parsed, err := Take(5)("Hello world")
//		fmt.Println(next)   // " world"
//		fmt.Println(parsed) // "Hello"
//		fmt.Println(err)    // nil
//	}
//
//	func failExample() {
//		next, parsed, err := Take(9999)("Hello world")
//		fmt.Println(next)   // ""
//		fmt.Println(parsed) // ""
//		fmt.Println(err)    // error
//	}
func Take(amount uint) Parser[string] {
	return func(input string) (string, string, error) {
		if amount > uint(len(input)) {
			return "", "", fmt.Errorf("the amount of characters to take is greater than the input size")
		}

		parsed := input[:amount]

		return input[amount:], parsed, nil
	}
}
