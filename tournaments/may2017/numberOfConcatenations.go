/**
Challenge: numberOfConcatenations
https://codefights.com/tournaments/vbZdzDTcajNBS7xN6/F

You have three arrays of integers: a, b, and c. Return the number of all triplets i, j, and k such that a[i]b[j] = c[k], where a[i]b[j] is the concatenation of two numbers in decimal notation.

Example

For a = [1, 2, 3, 4], b = [5, 6, 7, 8] and c = [17, 25, 83, 27], the output should be
numberOfConcatenations(a, b, c) = 3.

Here all the possible concatenations:

a[0]b[2] = 17 = c[0];
a[1]b[0] = 25 = c[1];
a[1]b[2] = 27 = c[3].
For a = [1, 1, 1], b = [1] and c = [11], the output should be
numberOfConcatenations(a, b, c) = 3.

You can concatenate every number from a with b[0] to get c[0].

Input/Output

[time limit] 4000ms (go)
[input] array.integer a

1 ≤ a.length ≤ 105,
1 ≤ a[i] ≤ 1000.

[input] array.integer b

Guaranteed constraints:
1 ≤ b.length ≤ 105,
1 ≤ b[i] ≤ 1000.

[input] array.integer c

Guaranteed constraints:
1 ≤ c.length ≤ 105,
1 ≤ c[i] ≤ 106.

[output] integer64

The number of all triplets i, j, and k such that a[i]b[j] = c[k].
**/
func numberOfConcatenations(a []int, b []int, c []int) int64 {

    mc := make(map[string]int)
    for _, cv := range c {
        cs := strconv.Itoa(cv)
        mc[cs]++
    }
    
    ma := make(map[string]int)
    mb := make(map[string]int)
    for _, av := range a {
        ma[strconv.Itoa(av)]++
    }
    for _, bv := range b {
        mb[strconv.Itoa(bv)]++
    }

    count := int64(0)
    for ak, av := range ma {
        for bk, bv := range mb {
            count += int64(mc[string(ak+bk)]*av*bv)
        }
    }
    return count
}
