/*
Challenge: Even Digits Only
https://codefights.com/arcade/intro/level-6/6cmcmszJQr6GQzRwW/description

Check if all digits of the given integer are even.

Example

For n = 248622, the output should be
evenDigitsOnly(n) = true;
For n = 642386, the output should be
evenDigitsOnly(n) = false.
Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
1 ≤ n ≤ 109.

[output] boolean

true if all digits of n are even, false otherwise.
*/
import "strings"
func evenDigitsOnly(n int) bool {
    s := strconv.Itoa(n)
    arr := strings.Split(s, "")
    for _, v := range arr {
        d, _ := strconv.Atoi(v)
        if d % 2 == 1 {
            return false
        }
    }
    return true
}
