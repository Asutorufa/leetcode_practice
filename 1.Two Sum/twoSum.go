package main

import "log"

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

var nums = []int{2, 7, 11, 15}
var target = 9

func one() {
	for x := range nums {
		for y := range nums {
			if nums[x] == 9-nums[y] {
				log.Println(x, y)
				return
			}
		}
	}
}

func main() {
	one()
}
