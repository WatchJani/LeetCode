package main

import (
	"fmt"
	"strings"
)

// func removeSpace(s string) string {
// 	for space := len(s) - 1; space >= 0; space-- {
// 		if s[space] != ' ' {
// 			return s[0 : space+1]
// 		}
// 	}

// 	return s
// }

// func lengthOfLastWord(s string) int {
// 	var counter int

// 	if s[len(s)-1] == ' ' {
// 		s = removeSpace(s)
// 	}

// 	fmt.Println(len(s))

// 	for index := len(s) - 1; index >= 0; index-- {
// 		if s[index] != ' ' {
// 			counter++
// 		} else {
// 			break
// 		}
// 	}

// 	return counter
// }

func lengthOfLastWord(s string) int {
	result := strings.Split(s, " ")

	for i := len(result) - 1; i >= 0; i-- {
		if len(strings.Trim(result[i], " ")) > 0 {
			return len(strings.Trim(result[i], " "))

		}

	}

	return 0
}

func main() {
	fmt.Println(lengthOfLastWord("a "))
}
