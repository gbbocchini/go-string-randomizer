package randomstring_test

import (
	"math/big"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/gbbocchini/go-randomstring"
)

func TestGenerateOne(t *testing.T) {
	r := randomstring.Randomizer{Universe: randomstring.LowerUpperDigits, Length: 13}
	s, err := r.GenerateOne()
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 13 {
		t.Fatalf("length = %d, want 13", len(s))
	}
	for _, c := range s {
		if !strings.ContainsRune(randomstring.LowerUpperDigits, c) {
			t.Fatalf("unexpected rune %q", c)
		}
	}
}

func TestGenerateOne_Deterministic(t *testing.T) {
	r1 := randomstring.Randomizer{Universe: randomstring.LowerUpperDigits, Length: 12, Seed: 42}
	r2 := randomstring.Randomizer{Universe: randomstring.LowerUpperDigits, Length: 12, Seed: 42}
	a, err := r1.GenerateOne()
	if err != nil {
		t.Fatal(err)
	}
	b, err := r2.GenerateOne()
	if err != nil {
		t.Fatal(err)
	}
	if a != b {
		t.Fatalf("expected deterministic output, got %q and %q", a, b)
	}
}

func TestGenerateOne_Unicode(t *testing.T) {
	const universe = "こんにちは世界"
	r := randomstring.Randomizer{Universe: universe, Length: 10}
	s, err := r.GenerateOne()
	if err != nil {
		t.Fatal(err)
	}
	if got := utf8.RuneCountInString(s); got != 10 {
		t.Fatalf("rune count = %d, want 10", got)
	}
	for _, c := range s {
		if !strings.ContainsRune(universe, c) {
			t.Fatalf("unexpected rune %q", c)
		}
	}
}

func TestGenerate(t *testing.T) {
	r := randomstring.Randomizer{Universe: randomstring.LowerLetters, Length: 6}
	got, err := r.Generate(100)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 100 {
		t.Fatalf("len = %d, want 100", len(got))
	}
	for _, s := range got {
		if len(s) != 6 {
			t.Fatalf("word length = %d, want 6", len(s))
		}
	}
}

func TestGenerate_Unique(t *testing.T) {
	r := randomstring.Randomizer{Universe: randomstring.LowerLetters, Length: 4, Unique: true}
	got, err := r.Generate(5000)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 5000 {
		t.Fatalf("len = %d, want 5000", len(got))
	}
	seen := make(map[string]struct{}, len(got))
	for _, s := range got {
		if _, dup := seen[s]; dup {
			t.Fatalf("duplicate string %q", s)
		}
		seen[s] = struct{}{}
	}
}

func TestGenerate_Deterministic(t *testing.T) {
	r1 := randomstring.Randomizer{Universe: randomstring.LowerUpperDigits, Length: 8, Seed: 7}
	r2 := randomstring.Randomizer{Universe: randomstring.LowerUpperDigits, Length: 8, Seed: 7}
	a, err := r1.Generate(10)
	if err != nil {
		t.Fatal(err)
	}
	b, err := r2.Generate(10)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(a, b) {
		t.Fatalf("expected deterministic slices, got %v and %v", a, b)
	}
}

func TestGenerate_UniqueExceedsPermutations(t *testing.T) {
	r := randomstring.Randomizer{Universe: "ab", Length: 2, Unique: true}
	if _, err := r.Generate(5); err == nil {
		t.Fatal("expected error when amount exceeds permutations")
	}
}

func TestGenerate_Errors(t *testing.T) {
	bad := []randomstring.Randomizer{
		{Universe: "", Length: 5},
		{Universe: randomstring.LowerLetters, Length: 0},
		{Universe: randomstring.LowerLetters, Length: -1},
	}
	for _, r := range bad {
		if _, err := r.GenerateOne(); err == nil {
			t.Errorf("expected error for %+v", r)
		}
	}
	if _, err := (randomstring.Randomizer{Universe: randomstring.LowerLetters, Length: 5}).Generate(-1); err == nil {
		t.Error("expected error for negative amount")
	}
}

func TestGenerate_Secure(t *testing.T) {
	r := randomstring.Randomizer{Universe: randomstring.LowerUpperDigitsSymbols, Length: 32, Secure: true}
	s, err := r.GenerateOne()
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 32 {
		t.Fatalf("len = %d, want 32", len(s))
	}
}

func TestUniquePermutations(t *testing.T) {
	r := randomstring.Randomizer{Universe: randomstring.LowerLetters, Length: 3}
	want := big.NewInt(26 * 26 * 26)
	if got := r.UniquePermutations(); got.Cmp(want) != 0 {
		t.Errorf("UniquePermutations = %s, want %s", got, want)
	}
}

func BenchmarkGenerate(b *testing.B) {
	r := randomstring.Randomizer{Universe: randomstring.LowerUpperDigits, Length: 16}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = r.GenerateOne()
	}
}
