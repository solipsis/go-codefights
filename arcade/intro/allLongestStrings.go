/*
Challenge: All Longest Strings
https://codefights.com/arcade/intro/level-3/fzsCQGYbxaEcTr2bL/description

Given an array of strings, return another array containing all of its longest strings.

Example

For inputArray = ["aba", "aa", "ad", "vcd", "aba"], the output should be
allLongestStrings(inputArray) = ["aba", "vcd", "aba"].

Input/Output

[time limit] 4000ms (go)
[input] array.string inputArray

A non-empty array.

Guaranteed constraints:
1 ≤ inputArray.length ≤ 10,
1 ≤ inputArray[i].length ≤ 10.

[output] array.string

Array of the longest strings, stored in the same order as in the inputArray.

*/

func allLongestStrings(arr []string) []string {
    max := 0
    for _, s := range arr {
        if len(s) > max {
            max = len(s)
        }
    }
    res := make([]string, 0)
    for _, s := range arr {
        if len(s) == max {
            res = append(res, s)
        }
    }
    return res
}
