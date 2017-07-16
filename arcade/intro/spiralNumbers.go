/*
Challenge: spiralNumbers
https://codefights.com/arcade/intro/level-12/uRWu6K8E7uLi3Qtvx/description

Construct a square matrix with a size N × N containing integers from 1 to N * N in a spiral order, starting from top-left and in clockwise direction.

Example

For n = 3, the output should be

spiralNumbers(n) = [[1, 2, 3],
                    [8, 9, 4],
                    [7, 6, 5]]
Input/Output

[time limit] 4000ms (go)
[input] integer n

Matrix size, a positive integer.

Guaranteed constraints:
3 ≤ n ≤ 10.

[output] array.array.integer
*/
func spiralNumbers(n int) [][]int {   
    row := n
    col := n
    
    r := make([][]int, row)
    for i := range r {
        r[i] = make([]int, col)
    }
    
    lwall := 0
    rwall := n
    uwall := 0
    dwall := n
    
    x := 0
    y := 0
     
    for i := 1; i <= n*n; {
        //right
        for  ;x < rwall; x++ {
            r[y][x] = i
            i++
        }
        x--
        y++
        //down
        for ;y < dwall; y++ {
            r[y][x] = i
            i++
        }
        y--
        x--
        //left
        for ;x >= lwall; x-- {
            r[y][x] = i
            i++
        }
        x++
        y--
        
        //up
        for ;y > uwall; y-- {
            r[y][x] = i
            i++
        }
        y++
        x++
        
        lwall++
        uwall++
        dwall--
        rwall--     
    }   
    return r
}
