/*
Challenge: Alphabetic Shift
https://codefights.com/arcade/intro/level-6/PWLT8GBrv9xXy4Dui/description

Given a string, replace each its character by the next one in the English alphabet (z would be replaced by a).

Example

For inputString = "crazy", the output should be
alphabeticShift(inputString) = "dsbaz".

Input/Output

[time limit] 4000ms (go)
[input] string inputString

Non-empty string consisting of lowercase English characters.

Guaranteed constraints:
1 ≤ inputString.length ≤ 10.

[output] string

The result string after replacing all of its characters.
*/
func alphabeticShift(s string) string {
    
    arr := []uint8(s)
    for n := range arr {
        arr[n]++
        if arr[n] > 122 {
            arr[n] = 97
        }
    }
    return string(arr)
}
