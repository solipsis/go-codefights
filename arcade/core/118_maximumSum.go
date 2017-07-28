/**
Challenge: Maximum Sum
https://codefights.com/arcade/code-arcade/sorting-outpost/64koZSDqndwYxFZj6/description


You are given an array of integers a. Range sum query is defined by a pair of non-negative integers l and r (l <= r). An output to a range sum query on the given array is the sum of all elements of a with indices from l to r inclusive.

Find an algorithm that given a list of range sum queries q can rearrange the array a in such a way that the total sum of all of the query outputs is maximized.

Example

For a = [2, 1, 2] and q = [[0, 1]], the output should be
maximumSum(a, q) = 4.

Input/Output

[time limit] 4000ms (go)
[input] array.integer a

An initial array.

Guaranteed constraints:
2 ≤ a.length ≤ 10,
1 ≤ a[i] ≤ 10.

[input] array.array.integer q

Array of range sum queries, each query is an array of two non-negative integers.

Guaranteed constraints:
1 ≤ q.length ≤ 10,
0 ≤ q[i][0] ≤ q[i][1] < a.length.

[output] integer

Maximum possible total sum of the given range sum query outputs.
**/
import "sort"
func maximumSum(a []int, q [][]int) int {
    
    sort.Ints(a)
    m := make([]int, len(a))
    for _, pair := range q {
        for i := pair[0]; i <= pair[1]; i++ {
            m[i]++
        }
    }
    sort.Ints(m)
    sum := 0
    for i := len(a)-1; i >= 0; i-- {
        sum += m[i] * a[i]
    }

    return sum
}


