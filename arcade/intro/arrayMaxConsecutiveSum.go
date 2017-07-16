/*
Challenge: arrayMaxConsecutiveSum
https://codefights.com/arcade/intro/level-8/Rqvw3daffNE7sT7d5/description

Given array of integers, find the maximal possible sum of some of its k consecutive elements.

Example

For inputArray = [2, 3, 5, 1, 6] and k = 2, the output should be
arrayMaxConsecutiveSum(inputArray, k) = 8.
All possible sums of 2 consecutive elements are:

2 + 3 = 5;
3 + 5 = 8;
5 + 1 = 6;
1 + 6 = 7.
Thus, the answer is 8.
Input/Output

[time limit] 4000ms (go)
[input] array.integer inputArray

Array of positive integers.

Guaranteed constraints:
3 ≤ inputArray.length ≤ 105,
1 ≤ inputArray[i] ≤ 1000.

[input] integer k

An integer (not greater than the length of inputArray).

Guaranteed constraints:
1 ≤ k ≤ inputArray.length.

[output] integer

The maximal possible sum.
*/
func arrayMaxConsecutiveSum(arr []int, k int) int {
    
    max := 0
    for start := 0; start < k; start++ {
        max += arr[start]
    }
    sum := max
    for i := k; i < len(arr); i++ {
        sum -= arr[i-k]
        sum += arr[i]
        if sum > max {
            max = sum
        }
    }
    return max
}
