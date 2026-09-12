package randomstring_test

import (
	"fmt"

	"github.com/gbbocchini/go-randomstring"
)

func ExampleRandomizer_GenerateOne() {
	r := randomstring.Randomizer{
		Universe: randomstring.LowerUpperDigits,
		Length:   13,
		Seed:     1,
	}
	id, err := r.GenerateOne()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	// Output: 9g14r5YgIsx9v
}

func ExampleRandomizer_Generate() {
	r := randomstring.Randomizer{
		Universe: randomstring.LowerUpperDigits,
		Length:   8,
		Unique:   true,
	}
	ids, err := r.Generate(1000)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(ids))
	// Output: 1000
}

func ExampleRandomizer_GenerateOne_secure() {
	r := randomstring.Randomizer{
		Universe: randomstring.LowerUpperDigitsSymbols,
		Length:   32,
		Secure:   true,
	}
	token, err := r.GenerateOne()
	if err != nil {
		panic(err)
	}
	fmt.Println(len(token))
	// Output: 32
}
