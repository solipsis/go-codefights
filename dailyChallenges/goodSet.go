/**
Challenge: goodSet
https://codefights.com/challenge/vQ5eLi5Mwatuq8oMJ/solutions/v4t39uNkRfuDWcFpN

A set of integers is called a good set if it contains 3 distinct elements a, b, and c, such that a + b = c.

Return false if the given set is a good set, and return true otherwise.

Example

For someSet = [1, 2, 3], the output should be
goodSet(someSet) = false.

For someSet = [1, 2, 4], the output should be
goodSet(someSet) = true.

Input/Output

[time limit] 4000ms (go)
[input] array.integer someSet

A sorted array of non-negative integers.

3 ≤ someSet.size() ≤ 5000,
0 ≤ someSet[i] ≤ 106.

[output] boolean

Return false if someSet is a good set, otherwise return true.
**/
func goodSet(someSet []int) (x bool) {
    m := make(map[int]bool)
    for _, v := range someSet {
        m[v] = !x
    }
    for k, _ := range m {
        for _, v := range someSet {
       
            if m[k + v] && k != v && (k+v) != k && (k+v) != v {
                return x
            }
        } 
    }
    return !x
}
