/**
Challenge: swapInPermutation
https://codefights.com/tournaments/ZS9S4h3HFKsuWRMnD/C

The permutation p of numbers 1, 2, ..., n is called parity-preserving if all the numbers on even positions (p[0], p[2], ...) are even and all the numbers on odd positions (p[1], p[3], ...) are odd. Given a permutation p, determine whether it's possible to swap no more than one pair of its elements to obtain a parity-preserving permutation.

Example

For p = [2, 1, 6, 5, 4, 3], the output should be
swapInPermutation(p) = true.

The permutation is already parity-preserving.

For p = [2, 4, 6, 5, 1, 3], the output should be
swapInPermutation(p) = true.

We can swap the numbers 4 and 1 to obtain a parity-preserving permutation.

For p = [3, 4, 6, 5, 1, 2], the output should be
swapInPermutation(p) = false.

More than one pair of elements would have to be swapped to obtain a parity-preserving permutation.

Input/Output

[time limit] 4000ms (go)
[input] array.integer p

An array of integers that represents a permutation of numbers 1, 2, ..., n.

Guaranteed constraints:
1 ≤ p.length ≤ 105.

[output] boolean

Return true if it's possible to swap no more than one pair of elements to obtain a parity-preserving permutation, otherwise return false.
**/
func swapInPermutation(p []int) bool {

    fail := 0   
    for n, v := range p {
        if n % 2 != v % 2 {
            fail++
        }
    }
    if fail > 2 || fail == 1 {
        return false
    }
    
    if fail <= 2 {
        
    }
    return true
}
