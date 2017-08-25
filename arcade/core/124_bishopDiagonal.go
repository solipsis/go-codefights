/**
Challenge: BishopDiagonal
https://codefights.com/arcade/code-arcade/chess-tavern/wGLCfzpdcA2o9kSpD/description

In the Land Of Chess, bishops don't really like each other. In fact, when two bishops happen to stand on the same diagonal, they immediately rush towards the opposite ends of that same diagonal.

Given the initial positions (in chess notation) of two bishops, bishop1 and bishop2, calculate their future positions. Keep in mind that bishops won't move unless they see each other along the same diagonal.

Example

For bishop1 = "d7" and bishop2 = "f5", the output should be
bishopDiagonal(bishop1, bishop2) = ["c8", "h3"].



For bishop1 = "d8" and bishop2 = "b5", the output should be
bishopDiagonal(bishop1, bishop2) = ["b5", "d8"].

The bishops don't belong to the same diagonal, so they don't move.


Input/Output

[time limit] 4000ms (go)
[input] string bishop1

Coordinates of the first bishop in chess notation.

[input] string bishop2

Coordinates of the second bishop in the same notation.

[output] array.string

Coordinates of the bishops in lexicographical order after they check the diagonals they stand on.
**/
import "sort"
func bishopDiagonal(b1 string, b2 string) []string {
    x1, y1 := parse(b1)
    x2, y2 := parse(b2)
    // if share diagonal
    if abs(x1-x2) == abs(y1-y2) {
        mx, my := 1, 1
        if x1 < x2 {
            mx = -1
        }
        if y1 < y2 {
            my = -1
        }
        // calculate final positions
        for x1 >= 1 && x1 < 7 && y1 >= 1 && y1 < 7 {
            x1 += mx
            y1 += my
        }
        for x2 >= 1 && x2 < 7 && y2 >= 1 && y2 < 7 {
            x2 += -mx
            y2 += -my
        }
    }
    s1 := string(rune(x1) + 'a') + string(rune(y1) + '1')
    s2 := string(rune(x2) + 'a') + string(rune(y2) + '1')
    res := []string{s1, s2}
    sort.Strings(res)
    return res
}

// absolute value
func abs(x int) int {
    if x < 0 { return -x }
    return x
}

// parse position
func parse(b string) (int, int) {
    c := int(b[0] - 'a')
    r := int(b[1] - '1')
    return c, r
}