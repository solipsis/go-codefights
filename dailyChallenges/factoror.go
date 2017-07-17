/**
Challenge: factoror
https://codefights.com/challenge/5hxyXFytS8z9CevDZ

For an integer n, we'll define an operation called factoror as follows:

First, find the prime factorization of n in the format a1p1 * a2p2 * ...;
Next, change each power operation into multiplication, and each multiplication operation into addition: a1 * p1 + a2 * p2 + ...;
Repeat the process until the value doesn't change anymore. This value is the factoror of n.
Given an array of numbers, find the factoror of each of them, and return the number of steps it took to calculate it.

Example

For numbers = [14], the output should be
factoror(numbers) = [4].

There's only one number in numbers, so only one factoror should be calculated. Here's how the factoror of 14 can be computed:

1) 14 = 2 * 7 => 2 + 7 = 9;
2) 9 = 3 ^ 2 => 3 * 2 = 6; 
3) 6 = 3 * 2 => 3 + 2 = 5; 
4) 5 = 5 => 5, hit the loop.
Thus, the answer is 4.

Input/Output

[time limit] 4000ms (go)
[input] array.integer numbers

An array of distinct positive integers.
0 ≤ numbers.Length ≤ 200,
2 ≤ input[i] ≤ 107.

[output] array.integer

An array of the same length as numbers, where the ith number is the number of steps it takes to calculate the factoror of numbers[i].
**/
func factoror(z []int) []int {
    
    res := make([]int, len(z))
    for n, i := range z {
        
        m := make(map[int]int)
        lastSum := -1

        steps := 0
        sum := i
        for sum != lastSum {
            steps++
            m = make(map[int]int)
            for x := 2; x * x <= i; x++ {
                for i % x == 0 {
                    i /= x
                    m[x]++
                }
            }
            if i > 1 { m[i]++ }
            lastSum = sum
            sum = 0
            for k, v := range m {
                sum += k * v
            }
            i = sum

        }
        res[n] = steps
    }
    return res  
}
