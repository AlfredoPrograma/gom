package gom

import (
	"fmt"
	"testing"
)

func TestAlt(t *testing.T) {
	tests := []ParserTestCase[ParsersList[string], string]{
		{
			name:  "successful parse",
			input: "foo bar baz",
			params: ParsersList[string]{
				Match("none"),
				Match("foo"),
			},
			want: ParseResult[string]{
				next:   " bar baz",
				parsed: "foo",
				err:    nil,
			},
		},
		{
			name:  "none branch match error",
			input: "foo bar baz",
			params: ParsersList[string]{
				Match("hello"),
				Match("dont match"),
				Match("never match"),
			},
			want: ParseResult[string]{
				next:   "",
				parsed: "",
				err:    fmt.Errorf("none branch match"),
			},
		},
	}

	ExecParserTestCases[ParsersList[string], string](t, Alt, tests)
}
