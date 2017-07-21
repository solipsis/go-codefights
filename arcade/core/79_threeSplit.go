/**
Challenge: Three Split
https://codefights.com/arcade/code-arcade/well-of-integration/QmK8kHTyKqh8xDoZk/description

You have a long strip of paper with integers written on it in a single line from left to right. You wish to cut the paper into exactly three pieces such that each piece contains at least one integer and the sum of the integers in each piece is the same. You cannot cut through a number, i.e. each initial number will unambiguously belong to one of the pieces after cutting. How many ways can you do it?

It is guaranteed that the sum of all elements in the array is divisible by 3.

Example

For a = [0, -1, 0, -1, 0, -1], the output should be
threeSplit(a) = 4.

Here are all possible ways:

[0, -1] [0, -1] [0, -1]
[0, -1] [0, -1, 0] [-1]
[0, -1, 0] [-1, 0] [-1]
[0, -1, 0] [-1] [0, -1]
Input/Output

[time limit] 4000ms (go)
[input] array.integer a

Guaranteed constraints:
5 ≤ a.length ≤ 104,
-108 ≤ a[i] ≤ 108.

[output] integer

It's guaranteed that for the given test cases the answer always fits signed 32-bit integer type.
**/
unc threeSplit(arr []int) int {

    sum := 0
    for _,x := range arr {
        sum += x
    }
    goal := sum / 3
    found := 0
    runningSum := 0
    for i := 0; i < len(arr)-2; i++ {
        runningSum += arr[i]
        if runningSum == goal {
            bSum := 0
            for j := i+1; j < len(arr)-1; j++ {
                bSum += arr[j]
                if bSum == goal {
                    cSum := 0
                    for k := j+1; k < len(arr); k++ {
                        cSum += arr[k]
                        if cSum == goal && k == len(arr)-1 {
                            fmt.Printf("i: %v j: %v k: %v \n", i, j, k)
                            found++
                        }
                    }
                }
            }
        }
    }
    return found   
}