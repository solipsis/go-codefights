/**
Challenge: Decipher
https://codefights.com/arcade/code-arcade/lab-of-transformations/dB9drnfWzpiWznESA/description

Consider the following ciphering algorithm:

For each character replace it with its code.
Concatenate all of the obtained numbers.
Given a ciphered string, return the initial one if it is known that it consists only of lowercase letters.

Note: here the character's code means its decimal ASCII code, the numerical representation of a character used by most modern programming languages.

Example

For cipher = "10197115121", the output should be
decipher(cipher) = "easy".

Explanation: charCode('e') = 101, charCode('a') = 97, charCode('s') = 115 and charCode('y') = 121.

Input/Output

[time limit] 4000ms (go)
[input] string cipher

A non-empty string which is guaranteed to be a cipher for some other string of lowercase letters.

Guaranteed constraints:
2 ≤ cipher.length ≤ 100.

[output] string
**/
func decipher(cipher string) string {

    r := ""
    for len(cipher) > 0 {
        n := ""
        if rune(cipher[0]) == '1' {
            n = cipher[0:3]
        } else {
            n = cipher[0:2]
        }
        cipher = cipher[len(n):]
        i, _ := strconv.Atoi(n)
        r += string(i)
    }
    return r
}
