/**
Challenge: leftover
https://codefights.com/challenge/QfFQ8aFHHry5Fsbu9

This is a reverse challenge, enjoy!

[time limit] 4000ms (go)
[input] string s

Guaranteed constraints:
1 ≤ s.length ≤ 50.

[output] integer
**/
func leftover(s string) (x int) { 
    for _, r := range s {
        x += int(r)
    }
      
    return x % len(s)
}
