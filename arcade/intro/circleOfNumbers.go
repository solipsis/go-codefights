/*
Challenge: circle of numbers
https://codefights.com/arcade/intro/level-7/vExYvcGnFsEYSt8nQ/description

Consider integer numbers from 0 to n - 1 written down along the circle in such a way that the distance between any two neighbouring numbers is equal (note that 0 and n - 1 are neighbouring, too).

Given n and firstNumber, find the number which is written in the radially opposite position to firstNumber.

Example

For n = 10 and firstNumber = 2, the output should be
circleOfNumbers(n, firstNumber) = 7.



Input/Output

[time limit] 4000ms (go)
[input] integer n

A positive even integer.

Guaranteed constraints:
4 ≤ n ≤ 20.

[input] integer firstNumber

Guaranteed constraints:
0 ≤ firstNumber ≤ n - 1.

[output] integer
*/
func circleOfNumbers(n int, f int) int {
    
    if f < n/2 {
        return f + n/2
    } else {
        return f - n/2
    }
}
