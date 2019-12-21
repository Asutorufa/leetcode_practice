package main

import "bytes"
import "log"

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	var out []byte
	var max int = 0
	var maxSubString []byte
	for x := range s {
		if bytes.IndexByte(out, s[x]) == -1 {
			out = append(out, s[x])
			// log.Println(string(out))
			if len(out) > max {
				max = len(out)
				maxSubString = out
			}
		} else {
			out = append(out[bytes.IndexByte(out, s[x])+1:], s[x])
			if len(out) > max {
				max = len(out)
				maxSubString = out
			}
		}
	}
	log.Println(string(maxSubString), max)
	return max
}

func main() {
	test1 := "abcabcbb"
	lengthOfLongestSubstring(test1)
	test2 := "bbbbb"
	lengthOfLongestSubstring(test2)
	test3 := "pwwkew"
	lengthOfLongestSubstring(test3)
	test4 := "aab"
	lengthOfLongestSubstring(test4)
	test5 := "dvdf"
	lengthOfLongestSubstring(test5)
}
