/*
Challenge: firstNotRepeatingCharacter
https://codefights.com/interview-practice/task/uX5iLwhc6L5ckSyNC/description

Note: Write a solution that only iterates over the string once and uses O(1) additional memory, since this is what you would be asked to do during a real interview.

Given a string s, find and return the first instance of a non-repeating character in it. If there is no such character, return '_'.

Example

For s = "abacabad", the output should be
firstNotRepeatingCharacter(s) = 'c'.

There are 2 non-repeating characters in the string: 'c' and 'd'. Return c since it appears in the string first.

For s = "abacabaabacaba", the output should be
firstNotRepeatingCharacter(s) = '_'.

There are no characters in this string that do not repeat.

Input/Output

[time limit] 4000ms (go)
[input] string s

A string that contains only lowercase English letters.

Guaranteed constraints:
1 ≤ s.length ≤ 105.

[output] char

The first non-repeating character in s, or '_' if there are no characters that do not repeat.
*/
func firstNotRepeatingCharacter(s string) string {
    
    // iterate over once tracking count and position of first appearance
    m := make(map[rune]*Node)
    for n, c := range s {
        if _, ok := m[c]; !ok {
            m[c] = &Node{n, 0}
        }
        (*m[c]).count++
    }
    
    // find chararter with count 1 and first appearance
    minIndex := len(s)
    for _, v := range m {
        if v.count == 1 && v.first <= minIndex {
            minIndex = v.first
        }
    }
    if minIndex >= len(s) {
        return "_"
    }
   
    return string(s[minIndex])
}

type Node struct {
    first, count int
}
