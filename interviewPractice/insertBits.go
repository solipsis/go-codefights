/**
Challenge: insertBits
https://codefights.com/interview-practice/task/sjRBSjPSPDAJgFX64/description


Given an integer n, replace its bits starting from the bit at position a to the bit at position b, inclusive, with the bits of integer k. Count from the least significant bit to the most significant bit, starting from 0.

Example

For n = 1024, a = 1, b = 6 and k = 17, the output should be
insertBits(n, a, b, k) = 1058.
n = 100 0000 00002, k = 1 00012, 1058 = 100 0010 00102.

For n = 11, a = 1, b = 2 and k = 2, the output should be
insertBits(n, a, b, k) = 13.
n = 10112, k = 102, 13 = 11012.

For n = 1, a = 3, b = 4 and k = 3, the output should be
insertBits(n, a, b, k) = 25.
n = 0 00012, k = 112, 25 = 1 10012.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
0 ≤ n < 231.

[input] integer a

Guaranteed constraints:
0 ≤ a ≤ 30.

[input] integer b

Guaranteed constraints:
a ≤ b ≤ 30.

[input] integer k

Guaranteed constraints:
0 ≤ k < 2B - A + 1.

[output] integer

Return n with the bits from a to b replaced with k's bits.
**/
func insertBits(n int, a int, b int, k int) int {

    mask := 0
    for i := 0; i <= b-a; i++ {
        mask += 1 << uint(i+a)
    }
    n &= ^mask // clear
    n |= k << uint(a) // or
    return n
}
