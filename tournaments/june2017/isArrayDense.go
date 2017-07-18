/**
Challenge: isArrayDense
https://codefights.com/tournaments/ZS9S4h3HFKsuWRMnD/A

John loves arrays. He calls an array a dense if there exist some numbers x and y such that a contains only consecutive numbers x, x + 1, x + 2, ..., y, in any order. His parents know that he loves arrays and gave him one for his birthday. Now John wants to check whether this array a is dense or not.

Example

For a = [3, 4, 2, 1], the output should be
isArrayDense(a) = true.

This array contains the consecutive numbers 1, 2, 3, and 4.

For a = [5, 3, 1, 2], the output should be
isArrayDense(a) = false.

Input/Output

[time limit] 4000ms (go)
[input] array.integer a

Guaranteed constraints:
1 ≤ a.length ≤ 1000,
-1000 ≤ a[i] ≤ 1000.

[output] boolean

Return true if a is dense, otherwise return false.
**/
func isArrayDense(a []int) bool {
    m := make(map[int]bool)
    min := 2000
    max := -2000
    for _, v := range a {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
        m[v] = true
    }
    
    fmt.Println(m);
    fmt.Println("min", min);
    fmt.Println("max", max)
    if (len(m) < len(a)) {return false}
    return len(m) == max-min+1
    
    
}
