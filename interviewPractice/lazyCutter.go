/**
Challenge: lazyCutter
https://codefights.com/interview-practice/task/SLajTY4nZjhx8Xuit/description


You're having a big party and you're serving a pizza as a main dish. You got really tired of cutting the pizza, so you decided to do it as efficiently as possible by using no more than n cuts.

Each cut is required to be a straight line, and there is no requirement that the pizza pieces be the same size.

Given n, the number of straight cuts you're willing to make, find the maximum number of pieces you can cut the pizza into.

Example

For n = 1, the output should be
lazyCutter(n) = 2.
With a single cut, we can only split the pizza into two pieces.

For n = 2, the output should be
lazyCutter(n) = 4.
With two cuts, we can make 3 pieces if the cuts don't intersect, or 4 pieces if they do. We're looking for the largest number of cuts, so 4 is the correct answer.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
1 ≤ n ≤ 104.

[output] integer

The maximum number of pieces you can cut the pizza into using n straight cuts.
**/
func lazyCutter(n int) int {
    s := 2
    z := 2
    for n > 1 {
        s += z
        z++
        n--
    }
    return s
}
