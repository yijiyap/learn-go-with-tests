package romannumerals

import "strings"

type RomanNumeral struct {
	Arabic int
	Roman  string
}

var allRomanNumerals = []RomanNumeral{
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

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Arabic {
			result.WriteString(numeral.Roman)
			arabic -= numeral.Arabic
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	var result int

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Roman) {
			result += numeral.Arabic
			roman = strings.TrimPrefix(roman, numeral.Roman)
		}
	}
	return result
}
