/*
Challenge: chessKnight
https://codefights.com/arcade/intro/level-11/pwRLrkrNpnsbgMndb/description

Given a position of a knight on the standard chessboard, find the number of different moves the knight can perform.

The knight can move to a square that is two squares horizontally and one square vertically, or two squares vertically and one square horizontally away from it. The complete move therefore looks like the letter L. Check out the image below to see all valid moves for a knight piece that is placed on one of the central squares.



Example

For cell = "a1", the output should be
chessKnight(cell) = 2.



For cell = "c2", the output should be
chessKnight(cell) = 6.



Input/Output

[time limit] 4000ms (go)
[input] string cell

String consisting of 2 letters - coordinates of the knight on an 8 × 8 chessboard in chess notation.

[output] integer
*/
func chessKnight(cell string) int {

    c, r := parse(cell)
    count := 0  
    // top 
    if (r < 7) {
        if (c > 1) {
            count++
        }
        if (c < 8) {
            count++
        }
    }
    if (r < 8) {
        if (c > 2) {
            count++
        }
        if (c < 7) {
            count++
        }
    }
    // bottom
    if (r > 2) {
        if (c > 1) {
            count++
        }
        if (c < 8) {
            count++
        }
    }
    if (r > 1) {
        if (c > 2) {
            count++
        }
        if (c < 7) {
            count++
        }
    }
    return count
}

func parse(s string) (int, int) {
    c := int(s[0]) - 96
    r := int(s[1])- 48

    return c, r
}
