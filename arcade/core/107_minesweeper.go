/*
Challenge: Minesweeper
https://codefights.com/arcade/code-arcade/waterfall-of-integration/ZMR5n7vJbexnLrgaM/description

In the popular Minesweeper game you have a board with some mines and those cells that don't contain a mine have a number in it that indicates the total number of mines in the neighboring cells. Starting off with some arrangement of mines we want to create a Minesweeper game setup.

Example

For

matrix = [[true, false, false],
          [false, true, false],
          [false, false, false]]
the output should be

minesweeper(matrix) = [[1, 2, 1],
                       [2, 1, 1],
                       [1, 1, 1]]       
Check out the image below for better understanding:



Input/Output

[time limit] 4000ms (go)
[input] array.array.boolean matrix

A non-empty rectangular matrix consisting of boolean values - true if the corresponding cell contains a mine, false otherwise.

Guaranteed constraints:
2 ≤ matrix.length ≤ 5,
2 ≤ matrix[0].length ≤ 5.

[output] array.array.integer

Rectangular matrix of the same size as matrix each cell of which contains an integer equal to the number of mines in the neighboring cells. Two cells are called neighboring if they share at least one corner.


*/
func minesweeper(m [][]bool) [][]int {
    
    res := make([][]int, len(m))
    for i := 0; i < len(m); i++ {
        res[i] = make([]int, len(m[0]))
    }
    
    for r := 0; r < len(m); r++ {
        for c := 0; c < len(m[0]); c++ {
            sum := 0
            // above
            if r > 0 {
                if m[r-1][c] {
                    sum++
                }
                if c > 0  && m[r-1][c-1] {
                    sum++
                }
                if c < len(m[0])-1 && m[r-1][c+1] {
                    sum++
                }
            }
            // middle
            if c > 0 && m[r][c-1] {
                sum++
            }
            if c < len(m[0])-1 && m[r][c+1] {
                sum++
            }
            // below
            if r < len(m)-1 {
                if m[r+1][c] {
                    sum++
                }
                if c > 0 && m[r+1][c-1] {
                    sum++
                }
                if c < len(m[0])-1 && m[r+1][c+1] {
                    sum++
                }
            }
            res[r][c] = sum
        }
    }
    return res
}
