package main

import "strings"

type RomanNumeral struct {
	// unsigned integer, which means they cannot be negative. We use this because Roman
	// numerals can not be negative.
	Value  uint16
	Symbol string
}

var allRomanNumberals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	// A Builder is used to efficiently build a string using Write methods.
	// It minimizes the memory copying
	var result strings.Builder

	for _, numeral := range allRomanNumberals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var arabic uint16 = 0

	for _, numeral := range allRomanNumberals {
		// HasPrefix(s, prefix) checks whether string s starts with prefix
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			// TrimPrefix(s, prefix) removes the prefix from s
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return arabic
}
