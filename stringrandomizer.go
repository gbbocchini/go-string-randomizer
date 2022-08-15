// Package go_string_randomizer Package provides small and optimized functions to generate
//random strings that can be used as passwords, control ids, etc.
package go_string_randomizer

import (
	"fmt"
	"math"
	"math/rand"
)

type StringRandomizer struct {
	LettersUniverse string
	GeneratedMaxLen int
	NoCollisions    bool
	RandSeed        int
}

// GenerateOne - generates 1 randomized string based on the struct initialization vars
func (s StringRandomizer) GenerateOne() string {
	rand.Seed(int64(s.RandSeed))
	id := make([]rune, s.GeneratedMaxLen)
	letters := []rune(s.LettersUniverse)
	for i := 0; i < s.GeneratedMaxLen; i++ {
		id[i] = letters[rand.Intn(len(letters))]
	}
	return string(id)
}

// GenerateBulk - generates amount arg of randomized string based on the struct initialization vars
func (s StringRandomizer) GenerateBulk(amount int) (error, []string) {
	var err error
	var result []string

	if s.NoCollisions {
		err = s.validatePermutationsAmount(amount)
	}

	if err != nil {
		return err, result
	}

	rand.Seed(int64(s.RandSeed))
	collisionControl := make(map[string]bool)
	letters := []rune(s.LettersUniverse)

	for len(result) != amount {
		id := make([]rune, s.GeneratedMaxLen)
		for i := 0; i < s.GeneratedMaxLen; i++ {
			id[i] = letters[rand.Intn(len(letters))]
		}

		if s.NoCollisions {
			if !s.isCollision(collisionControl, string(id)) {
				result = append(result, string(id))
			}
		} else {
			result = append(result, string(id))
		}
	}
	return err, result
}

// GetUniquePermutationsAmt Gets the max number of possible permutations given the struct initialization args
func (s StringRandomizer) GetUniquePermutationsAmt() float64 {
	return math.Pow(
		float64(len(s.LettersUniverse)),
		float64(s.GeneratedMaxLen),
	)
}

//helper function that checks collisions when attrib NoCollisions on the struct is set to true
func (s StringRandomizer) isCollision(controlMap map[string]bool, item string) bool {
	if _, exists := controlMap[item]; !exists {
		controlMap[item] = true
		return false
	}
	return true
}

//helper function that validates the max unique permutations amount
func (s StringRandomizer) validatePermutationsAmount(amountToGenerate int) error {
	var err error
	if float64(amountToGenerate) > s.GetUniquePermutationsAmt() {
		err = fmt.Errorf("the amount desired for generation (%v) is higher than the possible combinations (%v)",
			amountToGenerate, s.GetUniquePermutationsAmt())
	}
	return err
}
