/*
Challenge: differentSquares
https://codefights.com/arcade/intro/level-12/fQpfgxiY6aGiGHLtv/description

Given a rectangular matrix containing only digits, calculate the number of different 2 × 2 squares in it.

Example

For

matrix = [[1, 2, 1],
          [2, 2, 2],
          [2, 2, 2],
          [1, 2, 3],
          [2, 2, 1]]
the output should be
differentSquares(matrix) = 6.

Here are all 6 different 2 × 2 squares:

1 2
2 2
2 1
2 2
2 2
2 2
2 2
1 2
2 2
2 3
2 3
2 1
Input/Output

[time limit] 4000ms (go)
[input] array.array.integer matrix

Guaranteed constraints:
1 ≤ matrix.length ≤ 100,
1 ≤ matrix[i].length ≤ 100,
0 ≤ matrix[i][j] ≤ 9.

[output] integer

The number of different 2 × 2 squares in matrix.
*/
func differentSquares(matrix [][]int) int {

    m := make(map[string]bool)
    for i := 0; i < len(matrix)-1; i++ {
        for j := 0; j < len(matrix[0])-1; j++ {
            s := strconv.Itoa(matrix[i][j]) + strconv.Itoa(matrix[i+1][j]) + strconv.Itoa(matrix[i][j+1]) + strconv.Itoa(matrix[i+1][j+1])
            m[s] = true
        }
    }  
    return len(m)
}
