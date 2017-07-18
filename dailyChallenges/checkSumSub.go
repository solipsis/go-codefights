/**
Challenge: checkSumSub
https://codefights.com/challenge/oDm5up3SPpMZqpEmy

You are given an array of integers arr and a target k. Your task is to figure out whether you can pick a subarray of arr such that its elements add up to k.

Example

For arr = [0, 1, 2, 3, 4, 5] and k = 7, the answer should be
checkSum(arr, k) = true.

The sum of the elements of the subarray [3, 4] is exactly 7, so the answer is true.

Input/Output

[time limit] 4000ms (go)
[input] array.integer arr

Guaranteed constraints:
0 ≤ arr.length ≤ 105,
-5000 < arr[i] < 5000.

[input] integer k

Guaranteed constraints:
-109 < k < 109.

[output] boolean

Return true if some subarray of arr contains elements that add up to k, otherwise return false.
**/
func checkSumSub(a []int, k int) bool {

    s := 0
    m := make(map[int]int)
    for _, v := range a {
        s += v
        if s == k || m[s - k] > 0 {
            return 1>0
        } 
        
        m[s] = 1
    }
    return 1<0
}
