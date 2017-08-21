/**
Challenge: Sum34
https://codefights.com/challenge/X8MtaW3zRCTSkkCnA


Given n, you have to find out the sum of the digits of the number of numbers with length n which are formed using only digits 3 and 4.

Example
For n = 4, the output should be
sum34(n) = 7.
Desired numbers with length 4 are: 3333, 3334, 3343, 3433, 4333, 3344, 3434, 4334, 4343, 4433, 3443, 3444, 4344, 4434, 4443, 4444. So, there are 16 desired numbers, and the answer is 1 + 6 = 7.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
1 ≤ n ≤ 5 · 104.

[output] integer
**/
import "strings"
import "math/big"
func sum34(n int) (r int) {

    x := big.NewInt
   
    // 2^n then split the digits as strings
    s := x(2).Exp(x(2), x(int64(n)), nil)
    a := strings.Split(s.String(), "")
    
    for _, v := range a {
        i, _ := strconv.Atoi(v)
        r += i
    }
    return 
}
