package main

import "log"

import "container/list"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

func two() {
	one := list.New()
	one.PushBack(2)
	one.PushBack(4)
	one.PushBack(3)
	two := list.New()
	two.PushBack(5)
	two.PushBack(6)
	two.PushBack(4)
	last := list.New()
	for e, f, tmp := one.Front(), two.Front(), 0; ; {
		if e == nil && f == nil {
			break
		}
		a, b := 0, 0
		if e != nil {
			a = e.Value.(int)
			e = e.Next()
		}
		if f != nil {
			b = f.Value.(int)
			f = f.Next()
		}
		sum := tmp + a + b
		tmp = sum / 10
		last.PushBack(sum % 10)
		log.Println(a, b)
	}
	for e := last.Front(); e != nil; e = e.Next() {
		log.Println(e.Value)
	}
}

func main() {
	one()
	two()
}

func one() {
	var x = []int{2, 4, 3}
	var y = []int{5, 6, 4}
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
}
