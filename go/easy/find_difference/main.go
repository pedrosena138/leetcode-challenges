package main

import "fmt"

func findTheDifference(s string, t string) byte {
	var result byte
	var counter map[byte]uint8 = make(map[byte]uint8)

	for i := range s {
		counter[s[i]] -= 1
	}

	for i := range t {
		counter[t[i]] += 1
	}

	for k, v := range counter {
		if v == 1 {
			result = k
		}
	}

	return result
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}
