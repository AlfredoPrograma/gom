# gom

A simple and composable parser combinator library for Go inspired in [Nom](https://github.com/rust-bakery/nom) crate for Rust.

## Features

- **Composable Parsers:** Build complex parsers by combining simple ones.
- **Predicate-based Parsing:** Easily parse based on custom predicates.
- **Sequence, Choice, and Repetition:** Includes combinators for sequencing, alternatives, and repetition.

## Installation

```sh
go get github.com/alfredoprograma/gom
```

## Usage

### Basic Parsers

```go
import "github.com/alfredoprograma/gom"

next, parsed, err := gom.Match("foo")("foobar")
fmt.Println(parsed) // "foo"
fmt.Println(next)   // "bar"
```

### Combining Parsers

```go
// Sequence two parsers
pairParser := gom.Pair(gom.Match("foo"), gom.Match("bar"))
next, pair, err := pairParser("foobarbaz")
// pair.First == "foo", pair.Second == "bar", next == "baz"
```

### Predicate Parsers

```go
import "unicode"

letters := gom.TakeWhile(unicode.IsLetter)
next, parsed, err := letters("abc123")
// parsed == "abc", next == "123"
```

### Alternatives

```go
parsers := gom.ParsersList[string]{gom.Match("foo"), gom.Match("bar")}
altParser := gom.Alt(parsers)
next, parsed, err := altParser("barbaz")
// parsed == "bar", next == "baz"
```

## Testing

Run all tests:

```sh
make test
```

Check code coverage:

```sh
make coverage-terminal
```