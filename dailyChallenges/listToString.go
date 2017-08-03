/**
Challenge: listToString
https://codefights.com/challenge/iiuwDh9uf6AEEb4hz/solutions/y54noXuwSKK7hoJn9

This is a reverse challenge, enjoy!

[time limit] 4000ms (go)
[input] array.integer l

An array of integers with length 3.

Guaranteed constraints:
l.length = 3,
-105 ≤ l[i] ≤ 105.

[output] string
**/
func listToString(l []int) string {
    x := fmt.Sprintf
    a := x("%o", l[0])
    b := x("%b", l[1])
    if l[1] >= 0 {
        b = " " + b
    }
    c := x("%x", l[2])
    if l[2] >= 0 {
        c = " " + c
    }
  return a + b + c
}
