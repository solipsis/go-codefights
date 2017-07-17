/**
Challenge: Nav
https://codefights.com/challenge/CBiayArRfZEyDDGX8/solutions/kYKatj8xpJEhHMSmP

This is a reverse challenge, enjoy!

Input/Output

[time limit] 4000ms (go)
[input] string d

A string containing characters '>', '<', '+', '^', '-' and 'v'.

Guaranteed constraints:
1 ≤ d.length ≤ 1000.

[output] array.integer

An array of three elements.
**/
func Nav(d string) []int {

    var a,b,c int
    for _, r := range d {
        switch r {
            case '+':
            c++
            case '-':
            c--
            case '<':
            b--
            case '>':
            b++
            case 'v':
            a++
            case '^':
            a--
        }
    }
    return []int{a,b,c}
}

