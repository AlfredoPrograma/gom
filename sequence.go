package main

type PairResult[P1, P2 any] struct {
	first  P1
	second P2
}

func Pair[P1, P2 any](first Parser[P1], second Parser[P2]) Parser[PairResult[P1, P2]] {
	return func(input string) (string, PairResult[P1, P2], error) {
		parsed := PairResult[P1, P2]{}

		pipeNext, p1, err := first(input)

		if err != nil {
			return "", parsed, err
		}

		next, p2, err := second(pipeNext)

		if err != nil {
			return "", parsed, err
		}

		parsed.first = p1
		parsed.second = p2

		return next, parsed, nil
	}
}
