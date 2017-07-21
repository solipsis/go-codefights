/**
Challenge: singleNumber
https://codefights.com/interview-practice/task/APDXraJZYfPSYqQMJ/description


You are given an array of integers in which every element appears twice, except for one. Find the element that only appears one time. Your solution should have a linear runtime complexity (O(n)). Try to implement it without using extra memory.

Example

For nums = [2, 2, 1], the output should be
singleNumber(nums) = 1.

Input/Output

[time limit] 4000ms (go)
[input] array.integer nums

Guaranteed constraints:
1 ≤ nums.length < 104,
-109 ≤ nums[i] ≤ 109.

[output] integer
**/
func singleNumber(nums []int) int {
  res := 0
  for _, num := range nums {
    res ^= num
  }
  return res
}
