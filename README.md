# go-randomstring

> Fast, flexible random string generation for Go — ids, tokens, and passwords in one small, dependency-free package.

[![Go Reference](https://pkg.go.dev/badge/github.com/gbbocchini/go-randomstring.svg)](https://pkg.go.dev/github.com/gbbocchini/go-randomstring)
[![CI](https://github.com/gbbocchini/go-randomstring/actions/workflows/ci.yml/badge.svg)](https://github.com/gbbocchini/go-randomstring/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

- 🔤 **Flexible character sets** — seven built-in universes, or bring your own (Unicode included).
- 🔒 **Crypto-secure mode** — draw from `crypto/rand` for passwords and security tokens.
- 🔁 **Unique batches** — generate thousands of guaranteed-distinct strings in one call.
- 🎲 **Deterministic output** — pin a `Seed` for reproducible results in tests.
- 🚀 **Fast** — an ASCII fast path with zero dependencies beyond the standard library.

## Installation

```bash
go get github.com/gbbocchini/go-randomstring
```

## Quick start

```go
package main

import (
	"fmt"

	"github.com/gbbocchini/go-randomstring"
)

func main() {
	r := randomstring.Randomizer{
		Universe: randomstring.LowerUpperDigits,
		Length:   13,
	}
	id, err := r.GenerateOne()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
```

## Usage

### Basic generation

```go
r := randomstring.Randomizer{
	Universe: randomstring.LowerLetters,
	Length:   8,
}
id, err := r.GenerateOne()
```

### Unique batches

Set `Unique` to guarantee every string in a batch is distinct:

```go
r := randomstring.Randomizer{
	Universe: randomstring.LowerUpperDigits,
	Length:   8,
	Unique:   true,
}
ids, err := r.Generate(10_000) // 10,000 distinct ids
```

### Crypto-secure passwords

```go
r := randomstring.Randomizer{
	Universe: randomstring.LowerUpperDigitsSymbols,
	Length:   32,
	Secure:   true,
}
password, err := r.GenerateOne()
```

### Deterministic output (great for tests)

```go
r := randomstring.Randomizer{
	Universe: randomstring.LowerUpperDigits,
	Length:   13,
	Seed:     1,
}
id, _ := r.GenerateOne() // always "9g14r5YgIsx9v"
```

### Custom universe

```go
r := randomstring.Randomizer{
	Universe: "ABC123", // only these characters
	Length:   6,
}
```

### Built-in character sets

| Constant | Contents |
| --- | --- |
| `LowerLetters` | `a-z` |
| `UpperLetters` | `A-Z` |
| `Digits` | `0-9` |
| `Symbols` | `!@#$%&*()-_+={};:.,` |
| `LowerUpperLetters` | `a-z` + `A-Z` |
| `LowerUpperDigits` | `a-z` + `A-Z` + `0-9` |
| `LowerUpperDigitsSymbols` | `a-z` + `A-Z` + `0-9` + symbols |

## Documentation

Full API reference is available on [pkg.go.dev](https://pkg.go.dev/github.com/gbbocchini/go-randomstring).

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change, and update tests as appropriate.

## License

[MIT](LICENSE) © Gabriel Bocchini
