/*
Challenge: arrayMaximalAdjacentDifference
https://codefights.com/arcade/intro/level-5/EEJxjQ7oo7C5wAGjE/description


Given an array of integers, find the maximal absolute difference between any two of its adjacent elements.

Example

For inputArray = [2, 4, 1, 0], the output should be
arrayMaximalAdjacentDifference(inputArray) = 3.

Input/Output

[time limit] 4000ms (go)
[input] array.integer inputArray

Guaranteed constraints:
3 ≤ inputArray.length ≤ 10,
-15 ≤ inputArray[i] ≤ 15.

[output] integer

The maximal absolute difference.
*/
func arrayMaximalAdjacentDifference(arr []int) int {

    max := 0
    for i := 0; i < len(arr)-1; i++ {
        if abs(arr[i] - arr[i+1]) > max {
            max = abs(arr[i] - arr[i+1])
        }
    }
    return max
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
