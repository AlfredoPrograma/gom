package main

import (
	"reflect"
	"testing"
)

type ParseResult[O any] struct {
	next   string
	parsed O
	err    error
}

type ParserTestCase[P, O any] struct {
	name   string
	input  string
	params P
	want   ParseResult[O]
}

type ParserWrapper[P, O any] func(param P) Parser[O]

func ExecParserTestCases[P, O any](t *testing.T, parser ParserWrapper[P, O], tests []ParserTestCase[P, O]) {
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			next, parsed, err := parser(tc.params)(tc.input)
			got := ParseResult[O]{next, parsed, err}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s: expected %+v, but got %+v", tc.name, tc.want, got)
			}
		})

	}
}
