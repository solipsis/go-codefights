/**
Challenge: Chess Knight Moves
https://codefights.com/arcade/code-arcade/chess-tavern/zqDuSLMHhAknqnLtA/description


Given a position of a knight on the standard chessboard, find the number of different moves the knight can perform.

The knight can move to a square that is two squares horizontally and one square vertically, or two squares vertically and one square horizontally away from it. The complete move therefore looks like the letter L. Check out the image below to see all valid moves for a knight piece that is placed on one of the central squares.



Example

For cell = "a1", the output should be
chessKnightMoves(cell) = 2.



For cell = "c2", the output should be
chessKnightMoves(cell) = 6.



Input/Output

[time limit] 4000ms (go)
[input] string cell

String consisting of 2 letters - coordinates of the knight on an 8 × 8 chessboard in chess notation.

[output] integer
**/
func chessKnightMoves(cell string) (count int) {
    
    moves := [][]int{ {-2,1},{-1,2},{1,2},{2,1},{2,-1},{1,-2},{-1,-2},{-2,-1}}
    c := int(cell[0] - 'a')
    r := int(cell[1] - '1')  
    for _, m := range moves {
        x, y := m[0], m[1]
        if c + x >= 0 && c + x < 8 && r + y >= 0 && r + y < 8 {
            count++
        }
    }   
    return
}
