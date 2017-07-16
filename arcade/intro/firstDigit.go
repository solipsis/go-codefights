/*
Challenge: firstDigit
https://codefights.com/arcade/intro/level-8/rRGGbTtwZe2mA8Wov/description

Find the leftmost digit that occurs in a given string.

Example

For inputString = "var_1__Int", the output should be
firstDigit(inputString) = '1';
For inputString = "q2q-q", the output should be
firstDigit(inputString) = '2';
For inputString = "0ss", the output should be
firstDigit(inputString) = '0'.
Input/Output

[time limit] 4000ms (go)
[input] string inputString

A string containing at least one digit.

Guaranteed constraints:
3 ≤ inputString.length ≤ 10.

[output] char
*/
func firstDigit(inputString string) string {

    for _, r := range inputString {
        x := int(r)
        if x > 47 && x < 58 {
            return string(r)
        }
    }
    
    return ""
}
