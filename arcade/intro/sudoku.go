/*
Challenge: sudoku
https://codefights.com/arcade/intro/level-12/tQgasP8b62JBeirMS/description

Sudoku is a number-placement puzzle. The objective is to fill a 9 × 9 grid with digits so that each column, each row, and each of the nine 3 × 3 sub-grids that compose the grid contains all of the digits from 1 to 9.

This algorithm should check if the given grid of numbers represents a correct solution to Sudoku.

Example

For the first example below, the output should be true. For the other grid, the output should be false: each of the nine 3 × 3 sub-grids should contain all of the digits from 1 to 9.



Input/Output

[time limit] 4000ms (go)
[input] array.array.integer grid

A matrix representing 9 × 9 grid already filled with numbers from 1 to 9.

[output] boolean

true if the given grid represents a correct solution to Sudoku, false otherwise.
*/
func sudoku(grid [][]int) bool {

    for i := 0; i < 9; i++ {
        row := grid[i]
        m := make(map[int]bool)
        for _, num := range row {
            m[num] = true
        }
        if len(m) != 9 {
            return false
        }
        
    }
    
    for col := 0; col < 9; col++ {
        m := make(map[int]bool) 
        for j := 0; j < 9; j++ {
            m[grid[j][col]] = true
        }
        if len(m) != 9 {
            return false;
        }
        
    }
    
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {  
            m := make(map[int]bool)
            for x := 0; x < 3; x++ {
                for y := 0; y < 3; y++ {
                    m[grid[i*3 + x][j * 3 + y]] = true
                }
            }
            if len(m) != 9 {
                return false
            }      
        }  
    }
    return true
}

