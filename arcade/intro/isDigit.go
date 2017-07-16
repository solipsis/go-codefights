/*
Challenge: Is Digit
https://codefights.com/arcade/intro/level-11/Zr2XXEpkBPtYWqPJu/description

Determine if the given character is a digit or not.

Example

For symbol = '0', the output should be
isDigit(symbol) = true;
For symbol = '-', the output should be
isDigit(symbol) = false.
Input/Output

[time limit] 4000ms (go)
[input] char symbol

A character which is either a digit or not.

[output] boolean

true if symbol is a digit, false otherwise.
*/
func isDigit(symbol string) bool {
        _, err := strconv.Atoi(symbol)
        return err == nil
}
