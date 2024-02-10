package main

import "fmt"

type Predicate func(ch rune) bool

// Takes a predicate function and applies it sequentially over each character of the input string until evaluates to false.
//
// Returns the rest of the string, the accumulated characters which complied the predicate as string and a nil error.
//
// It is a flex parser. It means if parser doesn't match anything, it will not consume any character and will not return an error either; it will return input string, empty parsed string and nil error.
func FlexTakeWhile(predicate Predicate) Parser[string] {
	return func(input string) (string, string, error) {
		var parsed string

		for _, ch := range input {
			if !predicate(ch) {
				break
			}

			parsed += string(ch)
		}

		return input[len(parsed):], parsed, nil
	}
}

// Takes a predicate function and applies it sequentially over each character of the input string until evaluates to false.
//
// Returns the rest of the string, the accumulated characters which complied the predicate as string and a nil error.
//
// It is a strict parser. It means at least the first character must comply the predicate. If it doesn't, it will return empty values for next string and parsed string, and returns a fullfilled error.
func StrictTakeWhile(predicate Predicate) Parser[string] {
	return func(input string) (string, string, error) {
		var parsed string

		for _, ch := range input {
			if !predicate(ch) {
				break
			}

			parsed += string(ch)
		}

		if len(parsed) == 0 {
			return "", "", fmt.Errorf("at least one character should match with the predicate")
		}

		return input[len(parsed):], parsed, nil
	}
}

// Takes a predicate function and applies it sequentially over each character of the input string until evaluates to true.
//
// Returns the rest of the string, the accumulated characters which not complied the predacte as string and a nil error.
func FlexTakeTill(predicate Predicate) Parser[string] {
	return func(input string) (string, string, error) {
		var parsed string

		for _, ch := range input {
			if predicate(ch) {
				break
			}

			parsed += string(ch)
		}

		return input[len(parsed):], parsed, nil
	}
}
