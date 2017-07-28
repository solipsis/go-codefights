/**
Challenge: Sort By Length
https://codefights.com/arcade/code-arcade/sorting-outpost/QQB7f8ouAqY6jf7xi/description


Given an array of strings, sort them in the order of increasing lengths. If two strings have the same length, their relative order must be the same as in the initial array.

Example

For

inputArray = ["abc",
              "",
              "aaa",
              "a",
              "zz"]
the output should be

sortByLength(inputArray) = ["",
                            "a",
                            "zz",
                            "abc",
                            "aaa"]
Input/Output

[time limit] 4000ms (go)
[input] array.string inputArray

A non-empty array of strings.

Guaranteed constraints:
3 ≤ inputArray.length ≤ 10,
0 ≤ inputArray[i].length ≤ 10.

[output] array.string
**/
import "sort"
func sortByLength(arr []string) []string {

    bl := make(byLength, 0)
    for n, s := range arr {
        bl = append(bl, item{s, n})
    }
    sort.Sort(bl)
    res := make([]string, len(bl))
    for n, s := range bl {
        res[n] = s.s
    }
    return res
}

type item struct {
    s string
    pos int
}

type byLength []item
func (b byLength) Len() int { return len(b) }
func (b byLength) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byLength) Less(i, j int) bool {
    x, y := b[i], b[j]
    if len(x.s) < len(y.s) {
        return true
    }
    if len(x.s) == len(y.s) {
        if x.pos < y.pos {
            return true
        }
    }
    return false
}
