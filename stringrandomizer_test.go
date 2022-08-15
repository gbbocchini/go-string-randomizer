package go_string_randomizer

import "testing"

func TestGenerateOne(t *testing.T) {
	randomizer := StringRandomizer{
		LettersUniverse: LOWER_UPPER_LETTERS_NUMS,
		GeneratedMaxLen: 7,
		NoCollisions:    true,
		RandSeed:        123,
	}
	res := randomizer.GenerateOne()
	if len(res) != 7 {
		t.Errorf("Generated len should be 7")
	}
}

func TestGenerateBulk(t *testing.T) {
	randomizer := StringRandomizer{
		LettersUniverse: LOWER_UPPER_LETTERS_NUMS,
		GeneratedMaxLen: 10,
		NoCollisions:    true,
		RandSeed:        123,
	}
	err, res := randomizer.GenerateBulk(1000)

	if len(res) != 1000 && err == nil {
		t.Errorf("Generated slice len should be 1000")
	}

	if len(res[0]) != 10 && err == nil {
		t.Errorf("Generated word len should be 10")
	}
}

func TestIsCollision(t *testing.T) {
	randomizer := StringRandomizer{
		LettersUniverse: LOWER_UPPER_LETTERS_NUMS,
		GeneratedMaxLen: 10,
		NoCollisions:    true,
		RandSeed:        123,
	}
	ctrl := map[string]bool{"A": true, "B": false}
	res := randomizer.isCollision(ctrl, "A")

	if !res {
		t.Errorf("A should return true")
	}

	res = randomizer.isCollision(ctrl, "B")

	if !res {
		t.Errorf("B should return true")
	}
}
