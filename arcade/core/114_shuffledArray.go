/**
Challenge: Shuffled Array
https://codefights.com/arcade/code-arcade/sorting-outpost/s4BEFMcpLdGbjX9KX/description


A noob programmer was given two simple tasks: sum and sort the elements of the given array
a = [a1, a2, ..., an]. He started with summing and did it easily, but decided to store the sum he found in some random position of the original array which was a bad idea. Now he needs to cope with the second task, sorting the original array a, and it's giving him trouble since he modified it.

Given the array shuffled, consisting of elements a1, a2, ..., an, a1 + a2 + ... + an in random order, return the sorted array of original elements a1, a2, ..., an.

Example

For shuffled = [1, 12, 3, 6, 2], the output should be
shuffledArray(shuffled) = [1, 2, 3, 6].

1 + 3 + 6 + 2 = 12, which means that 1, 3, 6 and 2 are original elements of the array.

For shuffled = [1, -3, -5, 7, 2], the output should be
shuffledArray(shuffled) = [-5, -3, 2, 7].

Input/Output

[time limit] 4000ms (go)
[input] array.integer shuffled

Array of at least two integers. It is guaranteed that there is an index i such that shuffled[i] = shuffled[0] + ... + shuffled[i - 1] + shuffled[i + 1] + ... + shuffled[n].

Guaranteed constraints:
2 ≤ shuffled.length ≤ 15,
-300 ≤ shuffled[i] ≤ 300.

[output] array.integer

A sorted array of shuffled.length - 1 elements.
**/
import "sort"
func shuffledArray(shuffled []int) []int {

    m := make(map[int]bool)
    sum := 0
    
    for _, v := range shuffled {
        m[v] = true
        sum += v
    }
    
    // we look for a number that when subracted is a number in the array
    sumIdx := -1
    for n, v := range shuffled {
        if m[sum-v] {
            sumIdx = n
            break
        }
    }
    shuffled = append(shuffled[:sumIdx], shuffled[sumIdx+1:]...)
    sort.Ints(shuffled)
    return shuffled
}
