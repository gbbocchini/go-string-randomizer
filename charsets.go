package randomstring

// LowerLetters contains the lowercase ASCII letters a-z.
const LowerLetters = "abcdefghijklmnopqrstuvwxyz"

// UpperLetters contains the uppercase ASCII letters A-Z.
const UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Digits contains the decimal digits 0-9.
const Digits = "0123456789"

// Symbols contains a set of common punctuation and symbol characters.
const Symbols = "!@#$%&*()-_+={};:.,"

// LowerUpperLetters contains all lowercase and uppercase ASCII letters.
const LowerUpperLetters = LowerLetters + UpperLetters

// LowerUpperDigits contains all lowercase and uppercase ASCII letters plus
// digits.
const LowerUpperDigits = LowerLetters + UpperLetters + Digits

// LowerUpperDigitsSymbols contains all lowercase and uppercase ASCII letters,
// digits, and symbols.
const LowerUpperDigitsSymbols = LowerUpperDigits + Symbols
