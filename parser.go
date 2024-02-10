package main

import (
	"fmt"
	"strings"
)

// Represents the type of parser evaluation.
type ParserMode int

//   - Flex: flex parsers are parsers which will not fail if they dont match.
//
//   - Strict: strict parsers are parsers which will fail if they dont match at least one pattern.

const (
	FLEX ParserMode = iota
	STRICT
)

// Describes the signature function for parsers.
type Parser[O any] func(input string) (string, O, error)

// Takes a target rune and returns a parser which matches the first rune of the input string against the target.
//
// If it matches, returns the rest of the string, the stringified rune and a nil error.
// Else returns empty values for the next string and matched string, and returns a fullfilled error.
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
func Take(amount uint) Parser[string] {
	return func(input string) (string, string, error) {
		if amount > uint(len(input)) {
			return "", "", fmt.Errorf("the amount of characters to take is greater than the input size")
		}

		parsed := input[:amount]

		return input[amount:], parsed, nil
	}
}

// Takes a set of characters and returns a parser which tries to match the first character in the input string against one of the given characters.
//
// If it matches, returns the rest of the string, the character matched as string and a nil error.
// Else returns empty values for the next string and matched string, and returns a fullfilled error.
func OneOf(characters string) Parser[string] {
	return func(input string) (string, string, error) {
		parsed := string(input[0])

		for _, c := range characters {
			if parsed == string(c) {
				return input[1:], parsed, nil
			}
		}

		return "", "", fmt.Errorf("none given character matched")
	}
}

// Takes a set of characters and returns a parser which tries to match the first character in the input string against none of the given characters.
//
// If it doesnt match, returns the rest of the string, the first character as string and a nil error.
// Else returns empty values for the next string and matched string, and returns a fullfilled error.
func NoneOf(characters string) Parser[string] {
	return func(input string) (string, string, error) {
		firstChar := string(input[0])

		for _, c := range characters {
			if firstChar == string(c) {
				return "", "", fmt.Errorf("unexpected given character matched")
			}
		}

		return input[1:], firstChar, nil
	}
}

// Takes a target and returns a parser which tries to accumulate all the characters until reach the given target.
//
// If it matches, returns the rest input including the target, the accumulated characters before the target occurrence as string
// and a nil error.
func TakeUntil(target string) Parser[string] {
	return func(input string) (string, string, error) {
		parsed, next, found := strings.Cut(input, target)

		if !found {
			return input, "", nil
		}

		return target + next, parsed, nil
	}
}

// Same parsing proccess than [TakeUntil] but in a [STRICT] mode.
func StrictTakeUntil(target string) Parser[string] {
	return func(input string) (string, string, error) {
		parsed, next, found := strings.Cut(input, target)

		if !found {
			return "", "", fmt.Errorf("cannot match target")
		}

		return target + next, parsed, nil
	}
}
