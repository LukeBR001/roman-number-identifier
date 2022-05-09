package pkg

import (
	"errors"
	"identifier/pkg/model"
	"strings"
)

var validCombinations = []string{
	"CM",
	"CD",
	"XC",
	"XL",
	"IX",
	"IV",
}

var RomanNumbers = model.DefinedReferenceSymbols{
	Combinations: map[string]int{
		"I": 1,
		"V": 5, //não pode repetir
		"X": 10,
		"L": 50, //não pode repetir
		"C": 100,
		"D": 500, //não pode repetir
		"M": 1000,
	},
}

type Service interface {
	IdentifyCombinations(textPayload model.TextPayload) ([]string, error)
	IdentifyBiggerNumber(romanList []string) int
}

func IdentifyCombinations(textPayload model.TextPayload) ([]string, error) {

	identifiedCombinations := make([]string, 0)
	var str string

	for _, v := range textPayload.Text {
		vs := strings.ToUpper(string(v))
		if _, ok := RomanNumbers.Combinations[vs]; ok {
			str += vs
		} else {
			if str != "" {
				identifiedCombinations = append(identifiedCombinations, str)
			}
			str = ""
		}
	}

	if str != "" {
		identifiedCombinations = append(identifiedCombinations, str)
	}

	var validRomanNumbers []string
	for _, v := range identifiedCombinations {
		if isValidRomanSequence(v) {
			validRomanNumbers = append(validRomanNumbers, v)
		}
	}

	if len(validRomanNumbers) == 0 {
		return []string{}, errors.New("not found valid roman combinations")
	}

	return validRomanNumbers, nil
}

func isRepeatable(symbol string) bool {
	if symbol == "I" || symbol == "X" || symbol == "C" || symbol == "M" {
		return true
	}
	return false
}

func IdentifyBiggerNumber(romanList []string) int {
	var biggerGlobalNumber int

	for _, v := range romanList {
		var sumLocalNumbers int

		if len(v) > 1 {
			leftNumber := RomanNumbers.Combinations[string(v[0])]
			sumLocalNumbers = leftNumber

			for j := 1; j < len(v); j++ {

				actualString := string(v[j])
				actualNumber := RomanNumbers.Combinations[actualString]

				if actualNumber <= leftNumber {
					sumLocalNumbers += actualNumber

				} else {
					sumLocalNumbers = sumLocalNumbers + actualNumber - (2 * leftNumber)
				}
				leftNumber = RomanNumbers.Combinations[string(v[j])]
			}
			if sumLocalNumbers > biggerGlobalNumber {
				biggerGlobalNumber = sumLocalNumbers
			}

		} else {
			sumLocalNumbers = RomanNumbers.Combinations[string(v[0])]
			if sumLocalNumbers > biggerGlobalNumber {
				biggerGlobalNumber = RomanNumbers.Combinations[string(v[0])]
			}
		}
	}
	return biggerGlobalNumber
}

func isValidRomanSequence(s string) bool {
	if len(s) == 1 {
		return true
	}

	var check bool
	if len(s) > 1 {
		SymbolValue := RomanNumbers.Combinations
		leftSymbol := string(s[0])

		for i := 1; i < len(s); i++ {
			check = false
			str := s[:i]
			actualSymbol := string(s[i])

			if SymbolValue[leftSymbol] > SymbolValue[actualSymbol] {
				if strings.Contains(str, actualSymbol) {
					return false
				}
				check = true
			} else if SymbolValue[leftSymbol] == SymbolValue[actualSymbol] { // I == I
				count := strings.Count(str, actualSymbol)

				if isRepeatable(actualSymbol) {
					if count < 3 {
						check = true
					}
				}
			} else {
				if strings.Contains(str, actualSymbol) {
					return false
				}

				combSymbol := leftSymbol + actualSymbol
				for _, v := range validCombinations {
					if combSymbol == v {
						if lenStr := len(str); lenStr >= 2 {
							if SymbolValue[actualSymbol] > SymbolValue[string(str[lenStr-2])] {
								return false
							}
						}

						check = true
					}
				}
			}
			if !check {
				return check
			}
			leftSymbol = actualSymbol

		}
	}
	return check
}
