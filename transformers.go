package main

type MapFn[T, K any] func(parsed T) (K, error)

func Map[T, K any](parser Parser[T], mapper MapFn[T, K]) Parser[K] {
	return func(input string) (string, K, error) {
		var mapped K

		next, parsed, err := parser(input)

		if err != nil {
			return "", mapped, err
		}

		mapped, err = mapper(parsed)

		if err != nil {
			return "", mapped, err
		}

		return next, mapped, nil
	}
}
