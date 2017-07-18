/**
Challenge: totalBinSum
https://codefights.com/challenge/prd5XH6KJtWPRva4w

Given a number num in its binary representation, your task is to sum all the numbers in base 10 formed by the prefixes of num. More formally, sum up int(num[0, i]) for all possible i.

Since the answer can be very big, return it modulo 109 + 7.

Example

For bin = "1001", the output should be
totalBinSum(bin) = 16.

Here are all the prefixes:

12 = 110;
102 = 210;
1002 = 410;
10012 = 910.
Thus, the answer is 1 + 2 + 4 + 9 = 16.

Input/Output

[time limit] 4000ms (go)
[input] string num

A binary representation of some number, i.e. a string consisting of the characters '0' and '1'.

Guaranteed constraints:
2 ≤ num.length ≤ 5 · 104.

[output] integer

The sum of all the prefixes of num in base 10.
**/
import "math"

func totalBinSum(num string) int {

    sum := int64(0)
    prev := int64(0)
    mod := int64(math.Pow(10,9) + 7)
    for i := 0; i < len(num); i++ {
        val := (prev * 2) % mod
        if num[i] == '1' {
            val = (val + 1) % mod
        }
        prev = val
        sum = (sum + val) % mod
    }
    return int(sum)   
}
