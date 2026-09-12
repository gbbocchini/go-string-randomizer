package randomstring

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"unicode/utf8"

	mrand "math/rand/v2"
)

// Randomizer generates random strings according to its configuration.
//
// A Randomizer value may be reused and shared across goroutines: all fields
// are read only during generation, and no mutable state is kept on the value.
type Randomizer struct {
	// Universe is the set of runes from which generated strings are drawn.
	// It must not be empty. Multi-byte (Unicode) runes are supported.
	Universe string

	// Length is the number of runes in each generated string. It must be
	// greater than zero.
	Length int

	// Unique, when true, makes a single call to Generate return amount
	// distinct strings. It has no effect on GenerateOne.
	Unique bool

	// Secure, when true, draws randomness from crypto/rand, producing output
	// suitable for passwords and other security-sensitive tokens. Secure
	// generation ignores Seed.
	Secure bool

	// Seed makes output deterministic: two Randomizers with the same
	// configuration and the same non-zero Seed produce identical output.
	// The zero value (the default) seeds from a random source. Seed is
	// ignored when Secure is true.
	Seed int64
}

// GenerateOne generates a single random string.
//
// It returns an error if Universe is empty or Length is less than one.
func (r Randomizer) GenerateOne() (string, error) {
	if err := r.validate(); err != nil {
		return "", err
	}
	return r.generate(r.newPicker())
}

// Generate generates amount random strings.
//
// If Unique is set, the returned strings are all distinct. Generate returns an
// error if Universe is empty, Length is less than one, amount is negative, or
// Unique is set and amount exceeds the number of possible permutations.
func (r Randomizer) Generate(amount int) ([]string, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}
	if amount < 0 {
		return nil, fmt.Errorf("randomstring: amount must not be negative, got %d", amount)
	}

	pick := r.newPicker()

	if !r.Unique {
		out := make([]string, amount)
		for i := range out {
			s, err := r.generate(pick)
			if err != nil {
				return nil, err
			}
			out[i] = s
		}
		return out, nil
	}

	if max := r.UniquePermutations(); new(big.Int).SetInt64(int64(amount)).Cmp(max) > 0 {
		return nil, fmt.Errorf("randomstring: cannot generate %d unique strings from only %s permutations", amount, max)
	}

	out := make([]string, 0, amount)
	seen := make(map[string]struct{}, amount)
	for len(out) < amount {
		s, err := r.generate(pick)
		if err != nil {
			return nil, err
		}
		if _, dup := seen[s]; dup {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out, nil
}

// UniquePermutations returns the maximum number of distinct strings this
// Randomizer can produce, as a big.Int.
//
// It is the number of distinct runes in Universe raised to the power of
// Length. When Universe contains no duplicate runes this equals
// len(Universe)^Length.
func (r Randomizer) UniquePermutations() *big.Int {
	base := big.NewInt(int64(distinctRunes(r.Universe)))
	return new(big.Int).Exp(base, big.NewInt(int64(r.Length)), nil)
}

// validate reports whether the Randomizer is configured correctly.
func (r Randomizer) validate() error {
	if r.Universe == "" {
		return errors.New("randomstring: Universe must not be empty")
	}
	if r.Length < 1 {
		return fmt.Errorf("randomstring: Length must be greater than zero, got %d", r.Length)
	}
	return nil
}

// picker returns an index in [0, n).
type picker func(n int) (int, error)

// newPicker returns the random-index source for a single generation call,
// honoring Secure and Seed. The seeded source is constructed once so that a
// whole call advances through one deterministic sequence.
func (r Randomizer) newPicker() picker {
	switch {
	case r.Secure:
		return func(n int) (int, error) {
			v, err := crand.Int(crand.Reader, big.NewInt(int64(n)))
			if err != nil {
				return 0, err
			}
			return int(v.Int64()), nil
		}
	case r.Seed != 0:
		src := mrand.New(mrand.NewPCG(uint64(r.Seed), uint64(r.Seed)))
		return func(n int) (int, error) {
			return src.IntN(n), nil
		}
	default:
		return func(n int) (int, error) {
			return mrand.IntN(n), nil
		}
	}
}

// generate produces one random string using the given picker.
func (r Randomizer) generate(pick picker) (string, error) {
	if isASCII(r.Universe) {
		alphabet := []byte(r.Universe)
		buf := make([]byte, r.Length)
		for i := range buf {
			idx, err := pick(len(alphabet))
			if err != nil {
				return "", err
			}
			buf[i] = alphabet[idx]
		}
		return string(buf), nil
	}

	alphabet := []rune(r.Universe)
	buf := make([]rune, r.Length)
	for i := range buf {
		idx, err := pick(len(alphabet))
		if err != nil {
			return "", err
		}
		buf[i] = alphabet[idx]
	}
	return string(buf), nil
}

// isASCII reports whether s consists entirely of single-byte (ASCII) runes.
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

// distinctRunes reports the number of distinct runes in s.
func distinctRunes(s string) int {
	seen := make(map[rune]struct{})
	for _, r := range s {
		seen[r] = struct{}{}
	}
	return len(seen)
}
