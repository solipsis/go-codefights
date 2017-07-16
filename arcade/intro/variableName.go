/*
Challenge: variableName
https://codefights.com/arcade/intro/level-6/6Wv4WsrsMJ8Y2Fwno/description


Correct variable names consist only of Latin letters, digits and underscores and they can't start with a digit.

Check if the given string is a correct variable name.

Example

For name = "var_1__Int", the output should be
variableName(name) = true;
For name = "qq-q", the output should be
variableName(name) = false;
For name = "2w2", the output should be
variableName(name) = false.
Input/Output

[time limit] 4000ms (go)
[input] string name

Guaranteed constraints:
1 ≤ name.length ≤ 10.

[output] boolean

true if name is a correct variable name, false otherwise.
*/
import "regexp"
func variableName(name string) bool {
    r := regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
    return r.MatchString(name)
}
