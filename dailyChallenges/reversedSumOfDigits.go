package main

/*
Given integers p and n, find the smallest non-negative n-digit integer (represented as a string) whose digits add up to p. If there is no such number, return "-1" instead.

Example

For p = 15 and n = 3, the output should be
reversedSumOfDigits(p, n) = "159";

For p = 30 and n = 2, the output should be
reversedSumOfDigits(p, n) = "-1";

For p = 2 and n = 5, the output should be
reversedSumOfDigits(p, n) = "10001".

Input/Output

[execution time limit] 4 seconds (go)

[input] integer p

Guaranteed constraints:
0 ≤ p ≤ 105.

[input] integer n

Guaranteed constraints:
1 ≤ n ≤ 104.

[output] string
*/
import "strings"

func reversedSumOfDigits(p, n int) string {

	if p == 0 {
		if n == 1 {
			return "0"
		}
		return "-1"
	}

	d := []byte("1" + strings.Repeat("0", n-1))
	p -= 1

	n -= 1
	for {
		if n < 0 {
			return "-1"
		}
		if p >= 9 {
			p -= 9
			d[n] += 9
		} else {
			d[n] += byte(p)
			return string(d)
		}
		n -= 1
	}

}
