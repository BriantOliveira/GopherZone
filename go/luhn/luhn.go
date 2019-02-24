package luhn

import "unicode"

// Check if a credit card is valid
func Valid(input string) bool {
	readCard := false
	var allValues []int

	renes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		char := runes[i]

		if string(char) == " " {
			continue
		}

		if !unicode.IsNumber(char) {
			return false
		}

		newValue := int(char) - 48
		if readCard == true {
			newValue *= 2
			if newValue > 9 {
				newValue -= 9
			}
		}

		allValues = append(allValues, newValue)

		readCard = !readCard
	}
	if len(allValues) == 1 {
		return false
	}

	sum := 0
	for _, value := range allValues {
		sum += int(value)
	}

	if sum%10 == 0 {
		return true
	}
	return false
}
