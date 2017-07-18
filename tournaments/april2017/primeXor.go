/**
Challenge: primeXor
https://codefights.com/tournaments/BLhuiuSY4neuXPXet/E

The prime xor operation, a # b, takes two numbers and returns the product of the prime numbers that divide either a or b, but not both a and b.

Given two numbers a and b, return the result of operation a # b.

Example

For a = 18 and b = 10, the output should be
primeXor(a, b) = 15.

a = 2 * 3 * 3 and b = 2 * 5: Both factorizations contain the prime number 2, so 2 is not a candidate. 3 appears only in the factorization of a and 5 appears only in the factorization of b, so a # b = 3 * 5 = 15. Note that each prime number is counted once in a # b, even though 3 appears twice in the factorization of a.

Input/Output

[time limit] 4000ms (go)
[input] integer a

Guaranteed constraints:
2 ≤ a ≤ 107.

[input] integer b

Guaranteed constraints:
2 ≤ b ≤ 107.

[output] integer64

The result of operation a # b.
**/
func primeXor(a int, b int) int64 {

    am := factors(a)
    bm := factors(b)
    all := make(map[int]bool)
    
    for k, _ := range am {
        all[k] = true
    }
    for k, _ := range bm {
        all[k] = true
    }
    res := 1
    for k, _ := range all {
        if am[k] != bm[k] {
            res *= k
        }
    }
    return int64(res)
    
}

func factors(x int) map[int]bool {
    
    result := make(map[int]bool)
    for d := 2; d * d <= x; d++ {
        for x % d == 0 {
            result[d] = true
            x /= d
        }
    }
    if x > 1 {
        result[x] = true
    }
    
    return result
}
