/*
Challenge: differentSymbolsNaive
https://codefights.com/arcade/intro/level-8/8N7p3MqzGQg5vFJfZ/description

Given a string, find the number of different characters in it.

Example

For s = "cabca", the output should be
differentSymbolsNaive(s) = 3.

There are 3 different characters a, b and c.

Input/Output

[time limit] 4000ms (go)
[input] string s

A string of lowercase latin letters.

Guaranteed constraints:
3 ≤ s.length ≤ 15.

[output] integer
*/
func differentSymbolsNaive(s string) int {
    
    m := make(map[rune]bool)
    for _, c := range s {
        m[c] = 1 > 0
    }
    return len(m)
}
