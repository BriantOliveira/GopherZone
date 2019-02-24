package raindrops

import "strconv"

func Convert(num int) string {
	finalString := ""
	if num%3 == 0 {
		finalString += "Pling"
	}
	if num%5 == 0 {
		finalString += "Plang"
	}
	if num%7 == 0 {
		finalString += "Plong"
	}
	if finalString == "" {
		return strconv.Itoa(num)
	}
	return finalString
}