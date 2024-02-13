package main

import "fmt"

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
