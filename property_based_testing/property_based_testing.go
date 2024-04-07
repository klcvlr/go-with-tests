package property_based_testing

import (
	"strings"
)

type RomanNumeral struct {
	Value  uint16
	Symbol string
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

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder
	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var total uint16
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			total += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return total
}

func ConvertToArabicRecursive(roman string) uint16 {
	var convert func(roman string, index int) uint16
	convert = func(roman string, index int) uint16 {
		if index >= len(allRomanNumerals) {
			return 0
		}
		symbol := allRomanNumerals[index].Symbol
		value := allRomanNumerals[index].Value
		if strings.HasPrefix(roman, symbol) {
			roman = strings.TrimPrefix(roman, symbol)
			return value + convert(roman, index)
		}
		return convert(roman, index+1)
	}
	return convert(roman, 0)
}
