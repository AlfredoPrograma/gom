package main

import (
	"fmt"
)

type PairResult[T, K any] struct {
	first  T
	second K
}

func Pair[T, K any](firstParser Parser[T], secondParser Parser[K]) Parser[PairResult[T, K]] {
	return func(input string) (string, PairResult[T, K], error) {
		var result PairResult[T, K]
		rest, p1, err := firstParser(input)

		if err != nil {
			return "", result, fmt.Errorf("first parser failed")
		}

		next, p2, err := secondParser(rest)

		if err != nil {
			return "", result, fmt.Errorf("second parser failed")
		}

		result.first = p1
		result.second = p2

		return next, result, nil
	}
}

func Delimited[T, K, O any](opener Parser[T], parser Parser[O], closer Parser[K]) Parser[O] {
	return func(input string) (string, O, error) {
		next, _, err := opener(input)

		if err != nil {
			var parsed O
			return "", parsed, fmt.Errorf("opener parser failed")
		}

		next, parsed, err := parser(next)

		if err != nil {
			return "", parsed, fmt.Errorf("content parser failed")
		}

		next, _, err = closer(next)

		if err != nil {
			var parsed O
			return "", parsed, fmt.Errorf("closer parser failed")
		}

		return next, parsed, nil
	}
}

func Preceded[T, O any](preceded Parser[T], parser Parser[O]) Parser[O] {
	return func(input string) (string, O, error) {
		next, _, err := preceded(input)

		if err != nil {
			var parsed O
			return "", parsed, fmt.Errorf("preceded parser failed")
		}

		next, parsed, err := parser(next)

		if err != nil {
			var parsed O
			return "", parsed, fmt.Errorf("content parser failed")
		}

		return next, parsed, nil
	}
}
