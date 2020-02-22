package main

import (
	"log"
	"strings"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/

/*
my solution:
for example :
	Input: s = "PAYPALISHIRING", numRows = 4

	first covert the s to this:
	numRows + numRows - 2, so every row is six char
	0 1 2 3 4 5  <- index
	0 1 2 3 2 1  <- if the index > (numRows + numRows - 2) / 2, change the index to numRows + numRows - 2 - index
	P A Y P A L
	I S H I R I
	N G

	then just add the same index to one string:
	0: PIN
	1: ALSIG
	2: YAHR
	3: PI

	then add these string to one string:
	PINALSIGYAHRPI
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	str := []string{}
	tmp := strings.Builder{}
	for index := range s {
		tmp.WriteByte(s[index])
		if (index+1)%(numRows+numRows-2) == 0 {
			str = append(str, tmp.String())
			tmp.Reset()
		}
	}
	if tmp.String() != "" {
		str = append(str, tmp.String())
	}

	strSec := []string{}
	for i := 0; i < numRows; i++ {
		strSec = append(strSec, "")
	}
	for i := range str {
		for index := range str[i] {
			if index >= (numRows+numRows-2)/2 {
				// log.Println(index2, ((numRows+numRows-2)-index2)%numRows, string(str[index][index2]))
				strSec[(numRows+numRows-2)-index] += string(str[i][index])
			} else {
				// log.Println(index2-1, (index2)%numRows, string(str[index][index2]))
				strSec[index] += string(str[i][index])
			}
		}
	}

	tmp.Reset()
	for index := range strSec {
		tmp.WriteString(strSec[index])
	}

	// log.Println(strSec, tmp.String())
	return tmp.String()
}

/*
from leetcode solution
so easy ....
*/
func convertVisitByRow(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	str := strings.Builder{}
	cycleLen := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += cycleLen {
			// log.Println(j, i, j+i, j+cycleLen-i)
			str.WriteByte(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < len(s) {
				str.WriteByte(s[j+cycleLen-i])
			}
		}
	}
	return str.String()
}

func main() {
	log.Println(convert("PAYPALISHIRING", 1))
	log.Println(convert("PAYPALISHIRING", 2))
	log.Println(convert("PAYPALISHIRING", 3))
	log.Println(convert("PAYPALISHIRING", 4))
	log.Println(convert("PAYPALISHIRING", 5))
	log.Println(convertVisitByRow("PAYPALISHIRING", 1))
	log.Println(convertVisitByRow("PAYPALISHIRING", 2))
	log.Println(convertVisitByRow("PAYPALISHIRING", 3))
	log.Println(convertVisitByRow("PAYPALISHIRING", 4))
	log.Println(convertVisitByRow("PAYPALISHIRING", 5))
}
