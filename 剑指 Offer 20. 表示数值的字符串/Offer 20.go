package main

import (
	"fmt"
	"strings"
)

/**
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
例如，
字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，
但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
*/

/**
确定有限状态自动机
这题比较简单 与编译原理词法分析和语法分析有关
*/
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}
	if s == "." {
		return false
	}
	if strings.HasPrefix(s, ".e") {
		return false
	}
	lastSymbol := false
	first := true
	do := false
	lastDot := false
	e := false
	lastE := false

	for index := range s {
		if lastE && !lastDot && !lastSymbol {
			if symbol(s[index]) {
				first = false
				lastE = false
				lastSymbol = true
				continue
			}
		}
		if first {
			if symbol(s[index]) {
				first = false
				lastSymbol = true
				continue
			}
		}
		if !do && !e {
			if dot(s[index]) {
				first = false
				do = true
				lastDot = true
				continue
			}
		}
		if !first && !lastE && !e && !lastSymbol {
			if ee(s[index]) {
				e = true
				lastE = true
				lastDot = false
				continue
			}
		}
		if num(s[index]) {
			first = false
			lastE = false
			lastDot = false
			lastSymbol = false
			continue
		}
		return false
	}
	if lastE || lastSymbol {
		return false
	}
	return true
}

func symbol(b byte) bool {
	if b == '-' {
		return true
	}
	if b == '+' {
		return true
	}
	return false
}

func ee(b byte) bool {
	if b == 'E' || b == 'e' {
		return true
	}
	return false
}

func dot(b byte) bool {
	if b == '.' {
		return true
	}
	return false
}

func num(b byte) bool {
	if b == '0' {
		return true
	}
	if b == '1' {
		return true
	}
	if b == '2' {
		return true
	}
	if b == '3' {
		return true
	}
	if b == '4' {
		return true
	}
	if b == '5' {
		return true
	}
	if b == '6' {
		return true
	}
	if b == '7' {
		return true
	}
	if b == '8' {
		return true
	}
	if b == '9' {
		return true
	}
	return false
}

func main() {
	fmt.Println(isNumber("1.2"))        // true
	fmt.Println(isNumber("1e2"))        // true
	fmt.Println(isNumber("1e-2"))       // true
	fmt.Println(isNumber("1.2."))       // false
	fmt.Println(isNumber(".2"))         // true
	fmt.Println(isNumber("3."))         // true
	fmt.Println(isNumber("46.e3"))      // true
	fmt.Println(isNumber("aa"))         // false
	fmt.Println(isNumber("005047e+6"))  // true
	fmt.Println(isNumber("32.e-80123")) // true
	fmt.Println(isNumber("."))          // false
	fmt.Println(isNumber(".-4"))        // false
	fmt.Println(isNumber(".e1"))        // false
	fmt.Println(isNumber("+0e-"))       // false
	fmt.Println(isNumber("6e6.5"))      // false
}
