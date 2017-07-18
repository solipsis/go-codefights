/**
Challenge: nextGameOfLife
https://codefights.com/challenge/J9tdW6Fo7FiHmsqk4

Conway's Game of Life is a zero-player game that simulates a population of growing cells. Every cell interacts with its neighbors, which are the cells that are vertically, horizontally, or diagonally adjacent. At each step in time, the following changes take place:

Any live cell with fewer than two live neighbors dies (to simulate underpopulation);
Any live cell with two or three live neighbors lives on to the next generation;
Any live cell with more than three live neighbors dies (to simulate overpopulation);
Any dead cell with exactly three live neighbors becomes a live cell (to simulate reproduction).
In this challenge, the initial population is represented as a matrix of integers, seed, where 0 indicates a dead cell and 1 indicates a live cell. Your task here is to find the next population's pattern after one step in time.

Example:
For

seed = [[0, 1, 0], 
        [0, 1, 0], 
        [0, 1, 0]]
the output should be:

nextGameOfLife(seed) =
[[0, 0, 0], 
 [1, 1, 1], 
 [0, 0, 0]]
Visualisation: (credit to Wikipedia)


Input/Output

[time limit] 4000ms (go)
[input] array.array.integer seed

seed only contains arrays of 1's and 0's.

Guaranteed constraints:
1≤ seed.length ≤ 100,
1≤ seed[i].length ≤ 100,
All seed[i].length are guaranteed to be the same, no matter what i is.

[output] array.array.integer

Return what seed looks like after one step in time.
**/
func nextGameOfLife(s [][]int) [][]int {
    
    cp := make([][]int, len(s))
    for n := range cp {
        cp[n] = make([]int, len(s[n]))
        for nj, v := range s[n] {
            cp[n][nj] = v
        }
    }
    for ni := range s {
        for nj := range s[ni] {
            numNeighbors(ni, nj, s, cp)
        }
    }
    return cp
}

func numNeighbors(i, j int, s, cp [][]int) {
    c := 0
    // above
    if i > 0 {
        if s[i-1][j] > 0 {
            c++
        }
        if j > 0 && s[i-1][j-1] > 0 {
            c++
        }
        if j < len(s[0])-1 && s[i-1][j+1] > 0{
            c++
        }
    }
    // same row
    if j > 0 && s[i][j-1] > 0 {
        c++
    }
    if j < len(s[0])-1 && s[i][j+1] > 0{
        c++
    }
    // below
    if i < len(s)-1 {
        if s[i+1][j] > 0 {
            c++
        }
        if j > 0 && s[i+1][j-1] > 0 {
            c++
        }
        if j < len(s[0])-1 && s[i+1][j+1] > 0 {
            c++
        }
    }
    if s[i][j] > 0 && c < 2 {
        cp[i][j] = 0
    } else if s[i][j] > 0 && c < 4 {
        cp[i][j] = 1
    } else if s[i][j] == 1 && c > 3{
        cp[i][j] = 0
    } else if s[i][j] == 0 && c == 3 {
        cp[i][j] = 1
    } 
  
    
}