/**
Challenge: longestConsecutive
https://codefights.com/challenge/J9z6mzQQziyZuwPQJ

Given an array of non-negative integers arr, your task is to find and return the position of the first number in the array that has the longest number of consecutive set bits in its binary representation.

For instance, the longest number of consecutive bits in 11 is 2, since 1110 = 10112.

Example

For arr = [1, 2, 3, 4, 5, 6], the output should be
longestConsecutive(arr) = 2.

With all its values converted to binary, arr turns into ['1', '10', '11', '100', '101', '110']. As you can see, 3 and 6 have the longest streak of consecutively set bits (2). Since 3 comes before 6 in the array, and its position is 2, the output should be 2.

Input/Output

[time limit] 4000ms (go)
[input] array.integer arr

An array of non-negative integers.

Guaranteed constraints:
1 ≤ arr.length ≤ 105,
0 ≤ arr[i] < 231.

[output] integer

The position of the first number in the array that has the longest number of consecutive set bits in its binary representation.
**/
import "strings"

func longestConsecutive(a []int) int {
    
    m, i := 0,0
    for x, s := range a {
        for _, z := range strings.Split(strconv.FormatInt(int64(s), 2), "0") {
            if len(z) > m {
                m = len(z)
                i = x
            }
        }
    }
    
    return i
}