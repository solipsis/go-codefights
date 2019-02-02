/*
Challenge: ReverseInParentheses
https://app.codesignal.com/arcade/intro/level-3/9DgaPsE2a7M6M2Hu6/description?solutionId=fabgMpexGL9ryTJcr

Write a function that reverses characters in (possibly nested) parentheses in the input string.

Input strings will always be well-formed with matching ()s.

Example

For inputString = "(bar)", the output should be
reverseInParentheses(inputString) = "rab";
For inputString = "foo(bar)baz", the output should be
reverseInParentheses(inputString) = "foorabbaz";
For inputString = "foo(bar)baz(blim)", the output should be
reverseInParentheses(inputString) = "foorabbazmilb";
For inputString = "foo(bar(baz))blim", the output should be
reverseInParentheses(inputString) = "foobazrabblim".
Because "foo(bar(baz))blim" becomes "foo(barzab)blim" and then "foobazrabblim".
Input/Output

[execution time limit] 4 seconds (go)

[input] string inputString

A string consisting of lowercase English letters and the characters ( and ). It is guaranteed that all parentheses in inputString form a regular bracket sequence.

Guaranteed constraints:
0 ≤ inputString.length ≤ 50.

[output] string

Return inputString, with all the characters that were in parentheses reversed.

*/
import "strings"

func reverseInParentheses(s string) string {

	for strings.Contains(s, "(") {
		l := strings.LastIndex(s, "(")
		r := 0
		for n, ru := range s[l:] {
			if ru == ')' {
				r = n + l
				break
			}
		}

		s = s[:l] + reverse(s[l+1:r]) + s[r+1:]
	}
	return s
}

func reverse(s string) string {
	arr := []rune(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}
