/*
Challenge: chessboardCellColor
https://codefights.com/arcade/intro/level-6/t97bpjfrMDZH8GJhi/description

Given two cells on the standard chess board, determine whether they have the same color or not.

Example

For cell1 = "A1" and cell2 = "C3", the output should be
chessBoardCellColor(cell1, cell2) = true.



For cell1 = "A1" and cell2 = "H3", the output should be
chessBoardCellColor(cell1, cell2) = false.



Input/Output

[time limit] 4000ms (go)
[input] string cell1
[input] string cell2
[output] boolean

true if both cells have the same color, false otherwise.
*/
import "strings"
func chessBoardCellColor(cell1 string, cell2 string) bool {    
    return parse(cell1) == parse(cell2)
}

func parse(s string) bool {
    arr := strings.Split(s, "")
    a := uint8(arr[0][0]) - 65
    b := uint8(arr[1][0]) - 49
    return a % 2 == b % 2
}