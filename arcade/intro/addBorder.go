/*
Challenge: addBorder
https://codefights.com/arcade/intro/level-4/ZCD7NQnED724bJtjN/description

Given a rectangular matrix of characters, add a border of asterisks(*) to it.

Example

For

picture = ["abc",
           "ded"]
the output should be

addBorder(picture) = ["*****",
                      "*abc*",
                      "*ded*",
                      "*****"]
Input/Output

[time limit] 4000ms (go)
[input] array.string picture

A non-empty array of non-empty equal-length strings.

Guaranteed constraints:
1 ≤ picture.length ≤ 5,
1 ≤ picture[i].length ≤ 5.

[output] array.string

The same matrix of characters, framed with a border of asterisks of width 1.
*/
func addBorder(picture []string) []string {
    s := "*******"
    r := make([]string, 0)
    r = append(r, s[:len(picture[0])+2])
    for _, p := range picture {
        r = append(r, "*"+p+"*")
    }
    r = append(r, s[:len(picture[0])+2])
    return r
}
