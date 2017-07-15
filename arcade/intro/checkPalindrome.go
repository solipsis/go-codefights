/*
Challenge: checkPalindrome
https://codefights.com/arcade/intro/level-1/s5PbmwxfECC52PWyQ/description

Given the string, check if it is a palindrome.

Example

For inputString = "aabaa", the output should be
checkPalindrome(inputString) = true;
For inputString = "abac", the output should be
checkPalindrome(inputString) = false;
For inputString = "a", the output should be
checkPalindrome(inputString) = true.
Input/Output

[time limit] 4000ms (go)
[input] string inputString

A non-empty string consisting of lowercase characters.

Guaranteed constraints:
1 ≤ inputString.length ≤ 105.

[output] boolean

true if inputString is a palindrome, false otherwise.
*/

func checkPalindrome(inputString string) bool {
    return inputString == reverse(inputString)
}

func reverse(s string) string {
    arr := []rune(s)
    for i,j := 0, len(s)-1; i < j; i,j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }
    return string(arr)
}