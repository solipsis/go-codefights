/*
Challenge: longestDigitPrefix
https://codefights.com/arcade/intro/level-9/AACpNbZANCkhHWNs3/description

Given a string, output its longest prefix which contains only digits.

Example

For inputString="123aa1", the output should be
longestDigitsPrefix(inputString) = "123".

Input/Output

[time limit] 4000ms (go)
[input] string inputString

Guaranteed constraints:
3 ≤ inputString.length ≤ 35.

[output] string
*/
func longestDigitsPrefix(inputString string) string {

    var result bytes.Buffer
    for _, r := range inputString {
        if r >= 48 && r <= 57 {
            result.WriteRune(r)
        } else {
            break;
        }
    }
    
    return result.String()
}
