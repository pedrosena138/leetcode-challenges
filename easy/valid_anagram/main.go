package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var countS [26]int
	var countT [26]int

	for i := range s {
		countS[s[i]-'a']++
		countT[t[i]-'a']++

	}

	for i := range 26 {
		if countS[i] != countT[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
