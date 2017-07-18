/**
Challenge: swapPairBits
https://codefights.com/challenge/EYpoYCY6FNTKeimZx

Given a non-negative integer k, convert it to binary, swap the bits at positions p1 and p2 counting from the right, and return the obtained number in base 10.

Example

For k = 565, p1 = 9 and p2 = 3, the output should be
swapPairBits(k, p1, p2) = 817.

k2 = 1000110101 (Here, p1 is colored red and p2 is colored blue.) With p1 and p2 swapped, k becomes 1100110001, which is 81710. Thus, the final answer is 817.

Input/Output

[time limit] 4000ms (go)
[input] integer64 k

Guaranteed constraints
0 ≤ k ≤ 250.

[input] integer p1

Guaranteed constraints:
1 ≤ p1 ≤ 51.

[input] integer p2

Guaranteed constraints:
1 ≤ p2 ≤ 51.

[output] integer64

The number obtained by swapping the bits at p1 and p2.
**/
func swapPairBits(k int64, p1 int, p2 int) int64 {
    a := (k >> uint64(p1 - 1)) & 1
    b := (k >> uint64(p2 - 1)) & 1
    c := a ^ b
    d := (c << uint64(p1-1)) | (c << uint64(p2-1))
    return k ^ d
}
