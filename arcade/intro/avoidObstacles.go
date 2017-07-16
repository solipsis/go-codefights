/*
Challenge: avoidObstacles
https://codefights.com/arcade/intro/level-5/XC9Q2DhRRKQrfLhb5/description

You are given an array of integers representing coordinates of obstacles situated on a straight line.

Assume that you are jumping from the point with coordinate 0 to the right. You are allowed only to make jumps of the same length represented by some integer.

Find the minimal length of the jump enough to avoid all the obstacles.

Example

For inputArray = [5, 3, 6, 7, 9], the output should be
avoidObstacles(inputArray) = 4.

Check out the image below for better understanding:



Input/Output

[time limit] 4000ms (go)
[input] array.integer inputArray

Non-empty array of positive integers.

Guaranteed constraints:
2 ≤ inputArray.length ≤ 10,
1 ≤ inputArray[i] ≤ 40.

[output] integer

The desired length.
*/
func avoidObstacles(arr []int) int {
    m := make(map[int]bool)
    for _, i := range arr {
        m[i] = true
    }
    
    for l := 1; l < 50; l++ {
        con := false
        for x := 0; x <= 40; x += l {
            if m[x] {
                con = true
            }
        }
        if !con {
            return l
        }
    }
    return -1
}
