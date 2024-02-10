package main

import "fmt"

// Describes the signature for predicate function in predicate based parsers
type Predicate func(ch rune) bool

// Represents the compliance condition for the predicate.
const (
	COMPLY int = iota
	UNCOMPLY
)

// Helper function for define the predicate evaluation based on the parser mode and compliance condition.
func evalPredicate(input string, parserMode int, breakOn int, predicate Predicate) (string, error) {
	var accumulated string

	for _, ch := range input {
		if breakOn == COMPLY {
			if predicate(ch) {
				break
			}
		}

		if breakOn == UNCOMPLY {
			if !predicate(ch) {
				break
			}

		}

		accumulated += string(ch)
	}

	if parserMode == STRICT && len(accumulated) == 0 {
		return "", fmt.Errorf("at least one character should match with the predicate")
	}

	return accumulated, nil
}

// Takes a predicate function and applies it sequentially over each character of the input string until evaluates to false.
//
// Returns the rest of the string, the accumulated characters which complied the predicate as string and a nil error.
func TakeWhile(predicate Predicate) Parser[string] {
	return func(input string) (string, string, error) {
		parsed, _ := evalPredicate(input, FLEX, UNCOMPLY, predicate)

		return input[len(parsed):], parsed, nil
	}
}

// Takes a predicate function and applies it sequentially over each character of the input string until evaluates to false.
//
// Returns the rest of the string, the accumulated characters which complied the predicate as string and a nil error.
func StrictTakeWhile(predicate Predicate) Parser[string] {
	return func(input string) (string, string, error) {
		parsed, err := evalPredicate(input, STRICT, UNCOMPLY, predicate)

		if err != nil {
			return "", "", err
		}

		return input[len(parsed):], parsed, nil
	}
}

// Takes a predicate function and applies it sequentially over each character of the input string until evaluates to true.
//
// Returns the rest of the string, the accumulated characters which not complied the predacte as string and a nil error.
func TakeTill(predicate Predicate) Parser[string] {
	return func(input string) (string, string, error) {
		parsed, _ := evalPredicate(input, FLEX, COMPLY, predicate)

		return input[len(parsed):], parsed, nil
	}
}

// Takes a predicate function and applies it sequentially over each character of the input string until evaluates to true.
//
// Returns the rest of the string, the accumulated characters which not complied the predicate as string and a nil error.
func StrictTakeTill(predicate Predicate) Parser[string] {
	return func(input string) (string, string, error) {
		parsed, err := evalPredicate(input, STRICT, UNCOMPLY, predicate)

		if err != nil {
			return "", "", fmt.Errorf("at least one character should match with the predicate")
		}

		return input[len(parsed):], parsed, nil
	}
}
