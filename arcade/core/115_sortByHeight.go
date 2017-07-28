/**
Challenge: Sort by Height
https://codefights.com/arcade/code-arcade/sorting-outpost/D6qmdBL2NYz49XHwM/description


Some people are standing in a row in a park. There are trees between them which cannot be moved. Your task is to rearrange the people by their heights in a non-descending order without moving the trees.

Example

For a = [-1, 150, 190, 170, -1, -1, 160, 180], the output should be
sortByHeight(a) = [-1, 150, 160, 170, -1, -1, 180, 190].

Input/Output

[time limit] 4000ms (go)
[input] array.integer a

If a[i] = -1, then the ith position is occupied by a tree. Otherwise a[i] is the height of a person standing in the ith position.

Guaranteed constraints:
5 ≤ a.length ≤ 15,
-1 ≤ a[i] ≤ 200.

[output] array.integer

Sorted array a with all the trees untouched.
**/
import "sort"
func sortByHeight(a []int) []int {

    peeps := make([]int, 0)
    for _, i := range a {
        if i != -1 {
            peeps = append(peeps, i)
        }
    }
    sort.Ints(peeps)
    x := 0
    for n, i := range a {
        if i != -1 {
            a[n] = peeps[x]
            x++
        }
    }
    return a
}
