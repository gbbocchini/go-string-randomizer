// Package randomstring generates random strings for use as identifiers,
// tokens, passwords, and anywhere else you need short, configurable random
// text.
//
// The core type is Randomizer, configured through exported struct fields:
//
//	r := randomstring.Randomizer{
//		Universe: randomstring.LowerUpperDigits,
//		Length:   13,
//		Unique:   true,
//	}
//	id, err := r.GenerateOne()
//
// The package also provides ready-made character sets (LowerLetters,
// UpperLetters, Digits, Symbols, and their common combinations) that can be
// passed as a Randomizer's Universe.
package randomstring
