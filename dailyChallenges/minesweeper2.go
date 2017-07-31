/**
Challenge: minesweeper2
https://codefights.com/challenge/g6hd5k9Jfsfgs9hYF/


The second challenge of the minesweeper series!

As you click on a cell of the gameboard, there are 3 cases that thing will happen:

You clicked on a cell that contains a non-zero number (from 1 to 8). At that case, only that cell will open.
You clicked on a cell that contains a zero. At that case, 8 surrounding cells will also open. And if any of those 8 is a zero, its surrounding cell will continue to be opened.
You clicked on a mine (number 9). Boom, you lose immediately.
Given a gameboard consisting of either a 9 (a mine) or a number from 0 to 8 (the number of mines surrounding), a boolean opening matrix tells us whether a cell is opened or not, and an array of moves on that gameboard. Your task is to return a boolean matrix of the gameboard's state after these moves. If there is a move that makes the player loses, return an empty matrix.

Example
For

gameboard = [[0,1,9,1],
             [0,1,2,2],
             [0,0,1,9],
             [0,0,1,1]]
opening = [[false,false,false,false], 
           [false,false,false,false], 
           [false,false,false,false], 
           [false,false,false,false]]
and moves = [[3,0]], the output should be

minesweeper2(gameboard, opening, moves) =
[[true,true,false,false], 
 [true,true,true,false], 
 [true,true,true,false], 
 [true,true,true,false]]
The only move was made at the bottom-left corner. Because the value is 0, its surrounding cells continue to spread:


Input/Output

[time limit] 4000ms (go)
[input] array.array.integer gameboard

It's guaranteed that the gameboard is a valid gameboard.
(minesweeper1(gameboard) == true)
3 ≤ gameboard.length ≤ 50,
3 ≤ gameboard[i].length ≤ 50,
gameboard[i].length == gameboard[j].length,
0 ≤ gameboard[i][j] ≤ 9.

[input] array.array.boolean opening

It's guaranteed that there's no mine has been opened.
opening.length == gameboard.length,
opening[i].length == gameboard[i].length.

[input] array.array.integer moves

An array consisting of moves, where each move is a pair of numbers that indicate the cell's location. It's guaranteed that there's no cell was clicked twice.
0 ≤ moves.length ≤ 50,
moves[i].length == 2.

[output] array.array.boolean
**/
func minesweeper2(g [][]int, o [][]bool, m [][]int) [][]bool {
    for _, z := range m {       
        if g[z[0]][z[1]] == 9 {
            return make([][]bool, 0)
        }
        p(z[0], z[1], o, g)
    }
    return o
}

func p(r, c int, o [][]bool, g [][]int) {
    if r < 0 || r >= len(g) || c < 0 || c >= len(g[0]) || o[r][c] {
        return
    }   
    o[r][c] = true
    x := []int{-1,0,1}   
    if g[r][c] < 1 {
        for _, i := range x {
            for _, j := range x {
                p(r+i, c+j, o, g)
            }
        }
    }
}