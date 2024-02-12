package main

import "fmt"

type ParsersList[O any] []Parser[O]

func Alt[O any](parsers ParsersList[O]) Parser[O] {
	return func(input string) (string, O, error) {

		for _, p := range parsers {
			next, parsed, err := p(input)

			if err != nil {
				continue
			}

			return next, parsed, nil
		}

		var parsed O
		return "", parsed, fmt.Errorf("none branch match")
	}
}
