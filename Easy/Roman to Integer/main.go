package main

import "fmt"

// func romanToInt(s string) int {

// 	roman := map[string]int{
// 		"I":  1,
// 		"IV": 4,
// 		"V":  5,
// 		"IX": 9,
// 		"X":  10,
// 		"XL": 40,
// 		"L":  50,
// 		"XC": 90,
// 		"C":  100,
// 		"CD": 400,
// 		"D":  500,
// 		"CM": 900,
// 		"M":  1000,
// 	}

// 	var myNumber int

// 	if len(s) == 1 {
// 		return roman[string(s[0])]
// 	}

// 	for index := 0; index < len(s)-1; index++ {
// 		if value, check := roman[s[index:index+2]]; check {
// 			myNumber += value
// 			index += 1
// 		} else {
// 			myNumber += roman[string(s[index])]
// 		}
// 	}
// 	if _, check := roman[s[len(s)-2:]]; !check {
// 		myNumber += roman[string(s[len(s)-1])]
// 	}

// 	return myNumber
// }

var romanString = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {

	output := 0

	xFounded := false
	lFounded := false
	dFounded := false

	for i := len(s) - 1; i >= 0; i-- {

		val := string(s[i])

		if val == "X" || val == "V" {
			xFounded = true
		}

		if val == "L" || val == "C" {
			lFounded = true
		}

		if val == "D" || val == "M" {
			dFounded = true
		}

		if val == "I" {
			if xFounded {
				output -= getInt(val)
			} else {
				output += getInt(val)
			}
		} else if val == "X" {
			if lFounded {
				output -= getInt(val)
			} else {
				output += getInt(val)
			}
		} else if val == "C" {
			if dFounded {
				output -= getInt(val)
			} else {
				output += getInt(val)
			}
		} else {
			output += getInt(val)
		}
	}

	return output
}

func getInt(romanChar string) int {
	return romanString[romanChar]
}

func convertRomanToInt(s string) []int {
	integerStr := []int{}
	for _, c := range s {
		integerStr = append(integerStr, romanString[string(c)])
	}

	return integerStr
}

func main() {
	fmt.Println(romanToInt("D"))
}
