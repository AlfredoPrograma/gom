package main

import "fmt"

func evalRepetition[O any](input string, parser Parser[O], parserMode ParserMode) (string, []O, error) {
	accumulated := []O{}
	next := input

	for {
		n, p, e := parser(next)

		if e != nil || len(next) == 0 {
			break
		}

		accumulated = append(accumulated, p)
		next = n
	}

	if parserMode == STRICT && len(accumulated) == 0 {
		return "", accumulated, fmt.Errorf("parser should match at least one time")
	}

	return next, accumulated, nil
}

func Many[O any](parser Parser[O]) Parser[[]O] {
	return func(input string) (string, []O, error) {
		return evalRepetition[O](input, parser, FLEX)
	}
}

func StrictMany[O any](parser Parser[O]) Parser[[]O] {
	return func(input string) (string, []O, error) {
		return evalRepetition[O](input, parser, STRICT)
	}
}

func Count[O any](parser Parser[O], times uint) Parser[[]O] {
	return func(input string) (string, []O, error) {
		accumulated := []O{}
		next := input

		for i := 0; i < int(times); i++ {
			n, p, e := parser(next)

			if e != nil {
				return "", []O{}, fmt.Errorf("cannot execute parser %v times", times)
			}

			accumulated = append(accumulated, p)
			next = n
		}

		return next, accumulated, nil
	}
}
