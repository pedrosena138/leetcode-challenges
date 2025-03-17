package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1, word2 string) string {
	var strBuilder strings.Builder
	var maxLen int
	var word1Len int = len(word1)
	var word2Len int = len(word2)

	if word1Len > word2Len {
		maxLen = word1Len
	} else {
		maxLen = word2Len
	}

	for i := range maxLen {
		if i < word1Len {
			strBuilder.WriteString(string(word1[i]))
		}
		if i < word2Len {
			strBuilder.WriteString(string(word2[i]))
		}
	}

	return strBuilder.String()
}

func main() {
	fmt.Println(mergeAlternately("abc", "def"))
}
