/**
Challenge: bits
https://codefights.com/challenge/BiEcfpJAevtSMXAaQ

Given a array of positive integers, your task is to maximize each value in it. You have two ways to maximize a number:

Swap any two bits in its binary representation;
Do nothing. (Hey, you won't minimize a number this way either!)
Example

For numbers = [1, 5, 9], the output should be
bits(numbers) = [1, 6, 12].

Here's how each of the numbers can be maximized:

110 = 12, so there's nothing to be done with it;
510 = 1012, which can be maximized by swapping the last two bits to obtain 1102 = 610;
910 = 10012, which can be maximized by swapping the second and the fourth bits to obtain 11002 = 1210.
Input/Output

[time limit] 4000ms (go)
[input] array.integer64 numbers

Guaranteed constraints:
1 ≤ numbers.length ≤ 1000,
1 ≤ numbers[i] < 250.

[output] array.integer64

The maximized array of numbers.
**/
import "math"

func bits(z []int64) []int64 {

    // add leftmost 0 sub rightmost 1
    // if bigger replace
 
    for i, n := range z {
        // find leftmost zero
        s := int64(uint64(1) << uint64(math.Log2(float64(n))))
        for s > 0 {
            if s & n == 0 {
                break
            }
            s >>= 1
        }

        if x := n - n & -n + s; n < x {
            z[i] = x
        }
    }
    
    return z
}