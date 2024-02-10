package main

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
