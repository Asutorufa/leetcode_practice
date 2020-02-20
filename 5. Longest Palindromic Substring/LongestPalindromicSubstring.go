package main

import "log"

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	if len(s) == 1 || len(s) == 0 {
		return s
	}
	str, lenStr := string(s[0]), 0
	for i := 0; i < len(s); i++ {
		for x := len(s) - 1; x >= 0; x-- {
			if i >= x {
				break
			}
			if s[i] == s[x] {
				subStr, subStrLen := s[i:x+1], len(s[i:x+1])
				first, last := "", ""
				if subStrLen%2 == 0 {
					first, last = subStr[:subStrLen/2], subStr[subStrLen/2:]
				} else {
					first, last = subStr[:subStrLen/2], subStr[subStrLen/2+1:]
				}
				success := true
				for index := range first {
					if first[index] != last[len(last)-1-index] {
						success = false
					}
				}
				if success {
					if subStrLen > lenStr {
						lenStr = subStrLen
						str = subStr
					}
				}
			}
		}
	}
	return str
}

func main() {
	log.Println(longestPalindrome("abacab"))
	log.Println(longestPalindrome("babad"))
	log.Println(longestPalindrome("baabad"))
	log.Println(longestPalindrome("cbbd"))
	log.Println(longestPalindrome("aaaa"))
	log.Println(longestPalindrome("abadd"))

}
