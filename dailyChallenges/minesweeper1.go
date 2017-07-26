/**
Challenge: minesweeper1
https://codefights.com/challenge/ZXRv42mfhbiTTDrEB/


Let's try Minesweeper!
Each cell of Minesweeper gameboard can be:

a mine (appears as 9)
or a number representing the number of mines in its surrounding cells
(a cell is considered as surrounding another cell when this cell meets that cell on at least 1 corner) (appears as 0 - 8)
Your task is to check whether a gameboard is a valid gameboard.

Example
For

gameboard = [[0, 1, 9, 1],
             [0, 1, 1, 1],
             [0, 0, 0, 0]]
the output should be
minesweeper1(gameboard) = true.

Input/Output

[time limit] 4000ms (go)
[input] array.array.integer gameboard

Guaranteed constraints:
3 ≤ gameboard.length ≤ 100,
3 ≤ gameboard[i].length ≤ 100,
gameboard[i].length = gameboard[j].length,
0 ≤ gameboard[i][j] ≤ 9.

[output] boolean
**/
func minesweeper1(m [][]int) bool {

    res := make([][]int, len(m))
    for i := 0; i < len(m); i++ {
        res[i] = make([]int, len(m[0]))
    }
    
    for r := 0; r < len(m); r++ {
        for c := 0; c < len(m[0]); c++ {
            sum := 0
            // above
            if r > 0 {
                if m[r-1][c] == 9 {
                    sum++
                }
                if c > 0  && m[r-1][c-1] == 9 {
                    sum++
                }
                if c < len(m[0])-1 && m[r-1][c+1] == 9 {
                    sum++
                }
            }
            // middle
            if c > 0 && m[r][c-1] == 9 {
                sum++
            }
            if c < len(m[0])-1 && m[r][c+1] == 9 {
                sum++
            }
            // below
            if r < len(m)-1 {
                if m[r+1][c] == 9 {
                    sum++
                }
                if c > 0 && m[r+1][c-1] == 9 {
                    sum++
                }
                if c < len(m[0])-1 && m[r+1][c+1] == 9 {
                    sum++
                }
            }
            res[r][c] = sum
        }
    }
    
    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[0]); j++ {
            if m[i][j] != 9 && (m[i][j] != res[i][j]) {
                return false
            }
        }
    }
    
    
    return true
}