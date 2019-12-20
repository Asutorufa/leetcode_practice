package main

import "log"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

var x = []int{2, 4, 3}
var y = []int{5, 6, 4}

func one() []int {
	z := []int{}
	tmp := 0
	if len(x) < len(y) {
		tmp := x
		x = y
		y = tmp
	}
	for n := range x {
		sum := 0
		if n > len(y)-1 {
			sum = tmp + x[n] + 0
		} else {
			sum = tmp + x[n] + y[n]
		}
		tmp = sum / 10
		z = append(z, sum%10)
	}
	log.Println(z)
	return z
}

func main() {
	one()
}
