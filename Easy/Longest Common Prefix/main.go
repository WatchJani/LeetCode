package main

import "fmt"

//opet najbolji :D
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	} else {
		var myString string
		var counter int

		if len(strs[0]) == 0 {
			return ""
		}

		for {
			for index := 0; index < len(strs); index++ {
				if len(strs[index]) == counter {
					return myString
				}
				if !(strs[index][counter] == strs[0][counter]) {
					return myString
				}
			}
			myString += string(strs[0][counter])
			counter++
		}
	}
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"a", "ac"}))
}
