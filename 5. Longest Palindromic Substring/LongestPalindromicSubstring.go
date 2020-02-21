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

/*
 Brute Force
 more time,less memory
*/
func longestPalindromeBruteForce(s string) string {
	str, lenStr := "", 0
	for i := 0; i < len(s); i++ {
		for x := len(s) - 1; x >= 0; x-- {
			if i > x {
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

/*
 Expand Around Center
 less time,more memory
*/
func longestPalindromeExpandAroundCenter(s string) string {
	if s == "" {
		return ""
	}
	exandAroundCenter := func(s string, left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := exandAroundCenter(s, i, i)
		len2 := exandAroundCenter(s, i, i+1)
		lenS := 0
		if len1 > len2 {
			lenS = len1
		} else {
			lenS = len2
		}
		if lenS > end-start {
			start = i - (lenS-1)/2
			end = i + lenS/2
		}
	}
	return s[start : end+1]
}

func main() {
	log.Println(longestPalindromeBruteForce("abacab"))
	log.Println(longestPalindromeBruteForce("babad"))
	log.Println(longestPalindromeBruteForce("baabad"))
	log.Println(longestPalindromeBruteForce("cbbd"))
	log.Println(longestPalindromeBruteForce("aaaa"))
	log.Println(longestPalindromeBruteForce("abadd"))
	log.Println(longestPalindromeBruteForce(""))
	log.Println(longestPalindromeBruteForce("a"))
	log.Println(longestPalindromeExpandAroundCenter("abacab"))
	log.Println(longestPalindromeExpandAroundCenter("babad"))
	log.Println(longestPalindromeExpandAroundCenter("baabad"))
	log.Println(longestPalindromeExpandAroundCenter("cbbd"))
	log.Println(longestPalindromeExpandAroundCenter("aaaa"))
	log.Println(longestPalindromeExpandAroundCenter("abadd"))
	log.Println(longestPalindromeExpandAroundCenter(""))
	log.Println(longestPalindromeExpandAroundCenter("a"))

}
