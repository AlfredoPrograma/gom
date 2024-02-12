package main

import "fmt"

func main() {
	input := "foo bar baz"

	next, parsed, err := Alt[string](
		ParsersList[string]{
			Match("hazam"),
			Match("uwu"),
		},
	)(input)

	fmt.Println("[ERROR] ", err)
	fmt.Println("[NEXT] ", next)
	fmt.Println("[PARSED] ", parsed)
}
