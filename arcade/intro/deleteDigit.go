/*
Challenge: deleteDigit
https://codefights.com/arcade/intro/level-11/vpfeqDwGZSzYNm2uX/description

Given some integer, find the maximal number you can obtain by deleting exactly one digit of the given number.

Example

For n = 152, the output should be
deleteDigit(n) = 52;
For n = 1001, the output should be
deleteDigit(n) = 101.
Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
10 ≤ n ≤ 106.

[output] integer
*/
func deleteDigit(n int) int {

    max := 0
    s := strconv.Itoa(n)
    for n, _ := range s {
        temp := string(s[:n]) + string(s[n+1:])
        i, _ := strconv.Atoi(temp)
        if i > max {
            max = i
        }
    }
    return max
}
