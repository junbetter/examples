package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	var start, max int
	var hash = make(map[string]int)

	for i, v := range s {
		if _, ok := hash[string(v)]; ok {
			if hash[string(v)] > start {
				start = hash[string(v)]
			}
		}

		hash[string(v)] = i + 1
		curr := i - start + 1
		if curr > max {
			max = curr
		}
	}

	return max
}
