/**
Challenge: Rows Rearranging
https://codefights.com/arcade/code-arcade/sorting-outpost/vuXQuYFReJPe6hHAf/description


Given a rectangular matrix of integers, check if it is possible to rearrange its rows in such a way that all its columns become strictly increasing sequences (read from top to bottom).

Example

For

matrix = [[2, 7, 1], 
          [0, 2, 0], 
          [1, 3, 1]]
the output should be
rowsRearranging(matrix) = false;

For

matrix = [[6, 4], 
          [2, 2], 
          [4, 3]]
the output should be
rowsRearranging(matrix) = true.

Input/Output

[time limit] 4000ms (go)
[input] array.array.integer matrix

A 2-dimensional array of integers.

Guaranteed constraints:
1 ≤ matrix.length ≤ 10,
1 ≤ matrix[0].length ≤ 10,
-300 ≤ matrix[i][j] ≤ 300.

[output] boolean


**/
import "sort"
func rowsRearranging(matrix [][]int) bool {

    res := make(rows, 0)
    for _, row := range matrix {
        sort.Ints(row)
        res = append(res, row)
    }
    sort.Sort(res)
    for col := 0; col < len(matrix[0]); col++ {
        for row := 0; row < len(matrix)-1; row++ {
            if res[row][col] >= res[row+1][col] {
                return false
            }
        }
    }    
    return true
}

type rows [][]int
func (r rows) Len() int { return len(r) }
func (r rows) Swap(i, j int) { r[i], r[j] = r[j], r[i]}
func (r rows) Less(i, j int) bool {
    a, b := r[i], r[j]
    for i := 0; i < len(a); i++ {
        if a[i] < b[i] {
            return true
        }
        if b[i] < a[i] {
            return false
        }
        
    }
    return false
}