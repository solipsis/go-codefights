package main

/*
Remove all characters from a given string that appear more than once in it.

Example
For str = "zaabcbd", the output should be
removeDuplicateCharacters(str) = "zcd".

Input/Output

[execution time limit] 4 seconds (go)

[input] string str

A string consisting of English letters and spaces.

Guaranteed constraints:
5 ≤ str.length ≤ 30.

[output] string

Initial string where all characters appearing more than once are deleted.
*/

import z "strings"

func removeDuplicateCharacters(s string) (r string) {

	for _, c := range s {
		x := string(c)
		if len(z.Split(s, x)) < 3 {
			r += x
		}
	}
	return
}
