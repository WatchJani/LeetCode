package main

import (
	"fmt"
	"time"
)

func strStr(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	for i := 0; i < n-m+1; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}

func main() {
	new := time.Now()
	fmt.Println(strStr("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum~!", "~!"))
	fmt.Println(time.Since(new))
}
