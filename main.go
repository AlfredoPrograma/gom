package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func isNumber(ch rune) bool {
	return unicode.IsDigit(ch)
}

func mapToNumber(rawNumber string) (int, error) {
	return strconv.Atoi(rawNumber)
}

func main() {
	input := "Hello 123abcd world"

	next, parsed, _ := Pair[string, int](
		Match("Hello "),
		Map[string, int](
			FlexTakeWhile(isNumber),
			mapToNumber,
		),
	)(input)

	fmt.Println(next)
	fmt.Println(parsed)
}
