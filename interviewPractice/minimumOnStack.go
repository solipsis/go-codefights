/*
Challenge: minimumOnStack
https://codefights.com/interview-practice/task/gnZYGn367s4yaHvRr/description

Note: Write a solution with O(operations.length) complexity, since this is what you would be asked to do during a real interview.

Implement a modified stack that, in addition to using push and pop operations, allows you to find the current minimum element in the stack by using a min operation.

Example

For operations = ["push 10", "min", "push 5", "min", "push 8", "min", "pop", "min", "pop", "min"], the output should be
minimumOnStack(operations) = [10, 5, 5, 5, 10].

The operations array contains 5 instances of the min operation. The results array contains 5 numbers, each representing the minimum element in the stack at the moment when min was called.

Input/Output

[time limit] 4000ms (go)
[input] array.string operations

An array of operations consistently applied to the stack. For each valid i, operations[i] is one of the following operations:

push x: add the number x to the stack;
pop: remove an element from the top of the stack;
min: find the minimum element from the current stack elements.
It is guaranteed that pop and min are not applied to an empty stack. It is also guaranteed that all of the numbers in the stack are positive and not greater than 109.
Guaranteed constraints:
1 ≤ operations.length ≤ 104.

[output] array.integer

For each min operation, return the minimum element from the current stack elements at the moment when the operation was called.
*/
import "strings"

func minimumOnStack(operations []string) []int {

    m := MinStack{}
    m.stack = make([]int, 0)
    m.min = make([]int, 0)
    result := make([]int, 0)

    for _, s := range operations {
        args := strings.Split(s, " ")
        if s == "min" {
            result = append(result, m.Min())
        }
        if strings.Contains(s, "push") {
            i, _ := strconv.Atoi(args[1])
            m.Push(i)
        }
        if strings.Contains(s, "pop") {
            m.Pop()
        }
    }
    return result
}

type MinStack struct {
    stack []int
    min []int
}

func (s *MinStack) Push(i int) {
    if len(s.min) == 0 || i <= s.min[len(s.min)-1] {
        s.min = append(s.min, i)
    }
    s.stack = append(s.stack, i)
}

func (s *MinStack) Pop() int {
    
    if s.stack[len(s.stack)-1] == s.min[len(s.min)-1] {
        s.min = s.min[:len(s.min)-1]
    }
    val := s.stack[len(s.stack)-1]
    s.stack = s.stack[0:len(s.stack)-1]
    return val
}

func (s *MinStack) Min() int {
    return s.min[len(s.min)-1]
}