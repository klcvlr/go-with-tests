package property_based_testing

import (
	"testing"
	"testing/quick"
)

var cases = []struct {
	Description string
	arabic      uint16
	roman       string
}{
	{Description: "1 converts to I", arabic: 1, roman: "I"},
	{Description: "2 converts to II", arabic: 2, roman: "II"},
	{Description: "3 converts to III", arabic: 3, roman: "III"},
	{Description: "4 converts to IV", arabic: 4, roman: "IV"},
	{Description: "5 converts to V", arabic: 5, roman: "V"},
	{Description: "6 converts to VI", arabic: 6, roman: "VI"},
	{Description: "7 converts to VII", arabic: 7, roman: "VII"},
	{Description: "8 converts to VIII", arabic: 8, roman: "VIII"},
	{Description: "9 converts to IX", arabic: 9, roman: "IX"},
	{Description: "10 converts to X", arabic: 10, roman: "X"},
	{Description: "10 converts to X", arabic: 10, roman: "X"},
	{Description: "14 converts to XIV", arabic: 14, roman: "XIV"},
	{Description: "18 converts to XVIII", arabic: 18, roman: "XVIII"},
	{Description: "20 converts to XX", arabic: 20, roman: "XX"},
	{Description: "39 converts to XXXIX", arabic: 39, roman: "XXXIX"},
	{Description: "40 gets converted to XL", arabic: 40, roman: "XL"},
	{Description: "47 gets converted to XLVII", arabic: 47, roman: "XLVII"},
	{Description: "49 gets converted to XLIX", arabic: 49, roman: "XLIX"},
	{Description: "50 gets converted to L", arabic: 50, roman: "L"},
	{Description: "90 gets converted to XC", arabic: 90, roman: "XC"},
	{Description: "92 gets converted to XCII", arabic: 92, roman: "XCII"},
	{Description: "99 gets converted to XCIX", arabic: 99, roman: "XCIX"},
	{Description: "100 gets converted to C", arabic: 100, roman: "C"},
	{Description: "400 gets converted to CD", arabic: 400, roman: "CD"},
	{Description: "499 gets converted to CDXCIX", arabic: 499, roman: "CDXCIX"},
	{Description: "500 gets converted to D", arabic: 500, roman: "D"},
	{Description: "900 gets converted to CM", arabic: 900, roman: "CM"},
	{Description: "997 gets converted to CMXCVII", arabic: 997, roman: "CMXCVII"},
	{Description: "1000 gets converted to M", arabic: 1000, roman: "M"},
	{Description: "1984 gets converted to MCMLXXXIV", arabic: 1984, roman: "MCMLXXXIV"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			actual := ConvertToRoman(test.arabic)
			if actual != test.roman {
				t.Errorf("%s: roman: %q, actual: %q", test.Description, test.roman, actual)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			actual := ConvertToArabic(test.roman)

			if actual != test.arabic {
				t.Errorf("%s: roman: %d, actual: %d", test.Description, test.arabic, actual)
			}
		})
	}
}

func TestConvertToArabicRecursive(t *testing.T) {
	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			actual := ConvertToArabicRecursive(test.roman)

			if actual != test.arabic {
				t.Errorf("%s: roman: %d, actual: %d", test.Description, test.arabic, actual)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}

}

func BenchmarkConvertToArabic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range cases {
			actual := ConvertToArabic(test.roman)
			if actual != test.arabic {
				b.Errorf("%s: roman: %d, actual: %d", test.Description, test.arabic, actual)
			}
		}
	}
}

func BenchmarkConvertToArabicRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range cases {
			actual := ConvertToArabicRecursive(test.roman)
			if actual != test.arabic {
				b.Errorf("%s: roman: %d, actual: %d", test.Description, test.arabic, actual)
			}
		}
	}
}
