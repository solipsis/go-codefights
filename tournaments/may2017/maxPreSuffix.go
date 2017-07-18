/**
Challenge: maxPreSuffix
https://codefights.com/tournaments/vbZdzDTcajNBS7xN6/B

The PreSuffix of two numbers a and b is a number that is the prefix of one of the given numbers and the suffix of the other number. Both numbers are in decimal notation.

Given two integers, a and b, return their maximum PreSuffix, or -1 if they don't have any PreSuffixes.

Example

For a = 123456 and b = 456123, the output should be
maxPreSuffix(a, b) = 456.

In this case, a and b have two PreSuffixes, 123 and 456. 456 is the maximum PreSuffix.

For a = 123 and b = 456, the output should be
maxPreSuffix(a, b) = -1.

These two numbers don't have any PreSuffixes.

Input/Output

[time limit] 4000ms (go)
[input] integer a

Guaranteed constraints:
1 ≤ a ≤ 109.

[input] integer b

Guaranteed constraints:
1 ≤ b ≤ 109.

[output] integer

The maximum PreSuffix of a and b, or -1 if a PreSuffix doesn't exist.
**/
import "strings"
func maxPreSuffix(a int, b int) int {

    x := strconv.Itoa(a)
    y := strconv.Itoa(b)
    
    max := -1
    
    for i := 0; i < len(x); i++ {
        if strings.HasPrefix(y, string(x[i:]) ) {
            v, _ := strconv.Atoi(string(x[i:]))
            if v > max {
                max = v
            }
        }
    }
    
    for i := len(x)-1; i >= 0; i-- {
        if strings.HasPrefix(y, string(x[0:i+1]) ) {
            v, _ := strconv.Atoi(string(x[0:i+1]))
            if v > max {
                max = v
            }
        }
    }
    
    /////
    for i := 0; i < len(y); i++ {
        if strings.HasPrefix(x, string(y[i:]) ) {
            v, _ := strconv.Atoi(string(y[i:]))
            if v > max {
                max = v
            }
        }
    }
    
    for i := len(y)-1; i >= 0; i-- {
        if strings.HasPrefix(x, string(y[0:i+1]) ) {
            v, _ := strconv.Atoi(string(y[0:i+1]))
            if v > max {
                max = v
            }
        }
    }
    
    return max
}
