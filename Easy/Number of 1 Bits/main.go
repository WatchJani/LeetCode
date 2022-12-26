package main

import (
	"fmt"
	"strconv"
)

func hammingWeight(num uint32) int {
	t := strconv.FormatUint(uint64(num), 2)

	count := 0
	for _, v := range t {
		if v == '1' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(hammingWeight(120))
}
