/*
Challenge: adjacentElementsProduct
https://codefights.com/arcade/intro/level-2/xzKiBHjhoinnpdh6m/description

Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.

Example

For inputArray = [3, 6, -2, -5, 7, 3], the output should be
adjacentElementsProduct(inputArray) = 21.

7 and 3 produce the largest product.

Input/Output

[time limit] 4000ms (go)
[input] array.integer inputArray

An array of integers containing at least two elements.

Guaranteed constraints:
2 ≤ inputArray.length ≤ 10,
-1000 ≤ inputArray[i] ≤ 1000.

[output] integer

The largest product of adjacent elements.
*/

func adjacentElementsProduct(arr []int) int {

    max := -1001
    for i := 0; i < len(arr)-1; i++ {
        prod := arr[i] * arr[i+1]
        if prod > max {
            max = prod
        }
    }
    return max
}