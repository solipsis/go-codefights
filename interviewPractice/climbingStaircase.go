/**
Challenge: climbingStaircase
https://codefights.com/interview-practice/task/cAXEnPyzknC5zgd7x/description


You need to climb a staircase that has n steps, and you decide to get some extra exercise by jumping up the steps. You can cover at most k steps in a single jump. Return all the possible sequences of jumps that you could take to climb the staircase, sorted.

Example

For n = 4 and k = 2, the output should be

climbingStaircase(n, k) =
[[1, 1, 1, 1],
 [1, 1, 2],
 [1, 2, 1],
 [2, 1, 1],
 [2, 2]]
There are 4 steps in the staircase, and you can jump up 2 or fewer steps at a time. There are 5 potential sequences in which you jump up the stairs either 2 or 1 at a time.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
0 ≤ n ≤ 10.

[input] integer k

Guaranteed constraints:
0 ≤ k ≤ n.

[output] array.array.integer

An array containing all of the possible sequences in which you could climb n steps by taking them k or fewer at a time.
**/
func climbingStaircase(n int, k int) [][]int {
    
    res := make([][]int, 0)
    stack := make([]int, 0)
    
    if n == 0 {
        res = append(res, []int{})
        return res
    }
    
    stack = append(stack, 1)
    sum := 1
    for len(stack) > 0 {        
        if sum < n {
            stack = append(stack, 1)
            sum += 1
        }
        if sum == n {
            cp := make([]int, len(stack))
            copy(cp, stack)
            res = append(res, cp)
            stack, sum = pop(stack, sum) // pop
            // pop until we get to a number less than k
            for len(stack) > 0 && stack[len(stack)-1] >= k {
                stack, sum = pop(stack, sum)
            }
            if len(stack) > 0 {
                stack[len(stack)-1]++
                sum++
            }
        }
    }
    
    return res    
}

func pop(stack []int, sum int) ([]int, int) {
    sum -= stack[len(stack)-1]
    stack = stack[:len(stack)-1] // pop
    return stack, sum
}
