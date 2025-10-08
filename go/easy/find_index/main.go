package main

import "fmt"

func strStr(haystack string, needle string) int {
	var needleSize = len(needle)
	var maxSize = len(haystack)
	var haystackSlice []string

	var i int
	for i < maxSize {
		var pivot = i + needleSize
		if pivot <= maxSize {
			var sub_str = haystack[i : i+needleSize]
			haystackSlice = append(haystackSlice, sub_str)
		}
		i++
	}

	for c := range haystackSlice {
		if haystackSlice[c] == needle {
			return c
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
}
